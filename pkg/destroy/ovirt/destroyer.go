package ovirt

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	ovirtsdk "github.com/ovirt/go-ovirt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/openshift/installer/pkg/destroy/providers"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ovirt"
)

// ClusterUninstaller holds the various options for the cluster we want to delete.
type ClusterUninstaller struct {
	OvirtMetadata *ovirt.Metadata
	// InfraID
	InfraID string
	Logger  logrus.FieldLogger
}

// Run is the entrypoint to start the uninstall process.
func (o *ClusterUninstaller) Run() error {
	// - create ovirt connection from ovirt metadata
	cafile, err := os.Open(o.OvirtMetadata.Cafile)
	if cafile != nil {
		defer cafile.Close()
	}
	if err != nil {
		reader, err := getOvirtCA(o.OvirtMetadata.URL)
		cafile, err = ioutil.TempFile("", "ovirt-engine-ca.pem-*")
		if err != nil {
			return fmt.Errorf("failed to create ovirt-engine's ca file %s", err)
		}
		defer cafile.Close()
		_, err = io.Copy(cafile, reader)
		if err != nil {
			return fmt.Errorf("failed to write ovirt-engine ca's file %s", err)
		}
	}
	con, err := ovirtsdk.NewConnectionBuilder().
		URL(o.OvirtMetadata.URL).
		Username(o.OvirtMetadata.Username).
		Password(fmt.Sprint(o.OvirtMetadata.Password)).
		CAFile(cafile.Name()).
		Insecure(false).
		Build()

	if err != nil {
		return fmt.Errorf("failed to initialize connection to ovirt-engine's %s", err)
	}
	defer con.Close()

	// - find all vms by tag name=infraID
	vmsService := con.SystemService().VmsService()
	vmsResponse, err := vmsService.List().Search(fmt.Sprintf("tag=%s", o.InfraID)).Send()
	if err != nil {
		return err
	}
	// - stop + delete
	for _, vm := range vmsResponse.MustVms().Slice() {
		vmService := vmsService.VmService(vm.MustId())
		_, err = vmService.Stop().Send()
		o.Logger.Infof("Stopping VM %s : %s", vm.MustName(), errors.Wrapf(err, "success: %v", err != nil))
		_, err = vmService.Remove().Send()
		o.Logger.Infof("Removing VM %s : %s", vm.MustName(), errors.Wrapf(err, "success: %v", err != nil))
	}

	// finally remove the tag
	tagsService := con.SystemService().TagsService()
	tagsServiceListResponse, err := tagsService.List().Send()
	if err != nil {
		o.Logger.Warnf("Failed to fetch tags")
	}
	if tagsServiceListResponse != nil {
		for _, t := range tagsServiceListResponse.MustTags().Slice() {
			if t.MustName() == o.InfraID {
				_, err := tagsService.TagService(t.MustId()).Remove().Send()
				o.Logger.Infof("Removing tag %s : %s", t.MustName(), errors.Wrapf(err, "success: %v", err != nil))
			}
		}
	}
	return nil
}

// New returns oVirt Uninstaller from ClusterMetadata.
func New(logger logrus.FieldLogger, metadata *types.ClusterMetadata) (providers.Destroyer, error) {
	return &ClusterUninstaller{
		OvirtMetadata: metadata.ClusterPlatformMetadata.Ovirt,
		InfraID:       metadata.InfraID,
		Logger:        logger,
	}, nil
}

func getOvirtCA(engineURL string) (io.Reader, error) {
	u, err := url.Parse(engineURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ovirt-engine URL %s : %s", engineURL, err)
	}

	client := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	certURL := url.URL{
		Scheme:   u.Scheme,
		Host:     u.Host,
		Path:     "ovirt-engine/services/pki-resource",
		RawQuery: url.PathEscape("resource=ca-certificate&format=X509-PEM-CA"),
	}

	resp, err := client.Get(certURL.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error downloading ovirt-engine certificate %s with status %v", err, resp)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
