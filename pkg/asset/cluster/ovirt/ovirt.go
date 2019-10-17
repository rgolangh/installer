// Package ovirt extracts ovirt metadata from install configurations.
package ovirt

import (
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ovirt"
)

// Metadata converts an install configuration to ovirt metadata.
func Metadata(config *types.InstallConfig) *ovirt.Metadata {
	return &ovirt.Metadata{
		URL:        config.Platform.Ovirt.URL,
		Username:   config.Platform.Ovirt.Username,
		Password:   config.Platform.Ovirt.Password,
		Cafile:     config.Platform.Ovirt.Cafile,
		APIVIP:     config.Platform.Ovirt.APIVIP,
		DNSVIP:     config.Platform.Ovirt.DNSVIP,
		IngressVIP: config.Platform.Ovirt.IngressVIP,
		NodeMem:    config.Platform.Ovirt.NodeMem,
		NodeCores:  config.Platform.Ovirt.NodeCores,
	}
}
