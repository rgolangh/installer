package ovirt

// CloudProviderConfig represents ovirt's cloud provider config map
type CloudProviderConfig struct {
	URL               string `json:"ovirt_url,omitempty"`
	Username          string `json:"ovirt_username,omitempty"`
	Password          string `json:"ovirt_password,omitempty"`
	Cafile            string `json:"ovirt_cafile,omitempty"`
	Insecure          bool   `json:"ovirt_insecure,omitempty"`
	StorageDomainName string `json:"storage_domain_name,omitempty"`
	ClusterID         string `json:"cluster_id,omitempty"`
	TemplateID        string `json:"template_id,omitempty"`
}

// NewCloudProviderConfig creates a new CloudProviderConfig
func NewCloudProviderConfig(URL string, username string, password string, cafile string, insecure bool, storageDomainName string, clusterID string, templateID string) *CloudProviderConfig {
	return &CloudProviderConfig{
		URL:               URL,
		Username:          username,
		Password:          password,
		Cafile:            cafile,
		Insecure:          insecure,
		StorageDomainName: storageDomainName,
		ClusterID:         clusterID,
		TemplateID:        templateID}
}
