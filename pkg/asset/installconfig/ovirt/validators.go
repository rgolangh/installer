package ovirt

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	ovirtsdk "github.com/ovirt/go-ovirt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/openshift/installer/pkg/types/ovirt"
)

// Authenticated takes an ovirt platform and validates
// its connection to the API by establishing
// the connection and authenticating successfully.
// The API connection is closed in the end and must leak
// or be reused in any way.
func Authenticated(p *ovirt.Platform) survey.Validator {
	return func(val interface{}) error {
		engineURL, err := url.Parse(p.URL)
		if err != nil {
			return errors.Errorf("failed parse ovirt-engine survey URL %s", err)
		}

		if p.Cafile == "" {
			client := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
			certURL := url.URL{
				Scheme:   engineURL.Scheme,
				Host:     engineURL.Host,
				Path:     "ovirt-engine/services/pki-resource",
				RawQuery: url.PathEscape("resource=ca-certificate&format=X509-PEM-CA"),
			}

			logrus.Debugf("ovirt cert url %s", certURL.String())

			resp, err := client.Get(certURL.String())
			if err != nil || resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error downloading ovirt-engine certificate %s with status %v", err, resp)
			}
			defer resp.Body.Close()

			file, err := os.Create("/tmp/ovirt-engine.ca")
			if err != nil {
				return fmt.Errorf("failed writing ovirt-engine certificate %s", err)
			}
			io.Copy(file, resp.Body)
			p.Cafile = file.Name()
		}

		connection, err := ovirtsdk.NewConnectionBuilder().
			URL(p.URL).
			Username(p.Username).
			Password(fmt.Sprint(val)).
			CAFile(p.Cafile).
			Insecure(false).
			Build()

		if err != nil {
			return errors.Errorf("failed to construct connection to oVirt platform %s", err)
		}

		defer connection.Close()

		err = connection.Test()
		if err != nil {
			return errors.Errorf("connection to oVirt platform test failed %s", err)
		}
		return nil
	}

}
