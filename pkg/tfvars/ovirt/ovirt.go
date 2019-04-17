// Package ovirt contains ovirt-specific Terraform-variable logic.
package ovirt

import (
	"encoding/json"
)

type config struct {
	URL        string `json:"ovirt_url,omitempty"`
	Username   string `json:"ovirt_username,omitempty"`
	Password   string `json:"ovirt_password,omitempty"`
	Cafile     string `json:"ovirt_cafile,omitempty"`
	ClusterID  string `json:"ovirt_cluster_id,omitempty"`
	TemplateID string `json:"ovirt_template_id,omitempty"`
}

// TFVars generates ovirt-specific Terraform variables.
func TFVars(
	engineURL string,
	engineUser string,
	enginePass string,
	engineCafile string,
	clusterID string,
	templateID string) ([]byte, error) {

	cfg := config{
		URL:        engineURL,
		Username:   engineUser,
		Password:   enginePass,
		Cafile:     engineCafile,
		ClusterID:  clusterID,
		TemplateID: templateID,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
