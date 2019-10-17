package ovirt

// Platform stores all the global configuration that all
// machinesets use.
type Platform struct {
	URL               string `json:"ovirt_url,omitempty"`
	Username          string `json:"ovirt_username,omitempty"`
	Password          string `json:"ovirt_password,omitempty"`
	Cafile            string `json:"ovirt_cafile,omitempty"`
	Insecure          bool   `json:"ovirt_insecure,omitempty"`
	StorageDomainName string `json:"storage_domain_name,omitempty"`
	ClusterID         string `json:"cluster_id,omitempty"`
	TemplateID        string `json:"template_id,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on ovirt for machine pools which do not define their
	// own platform configuration.
	// +optional
	// Default will set the image field to the latest RHCOS image.
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	// APIVIP is an IP which will be served by bootstrap and then pivoted masters, using keepalived
	APIVIP string `json:"api_vip,omitempty"`

	// DNSVIP is the IP of the internal DNS which will be operated by the cluster
	DNSVIP string `json:"dns_vip,omitempty"`

	// ingressIP is an external IP which routes to the default ingress controller.
	// The IP is a suitable target of a wildcard DNS record used to resolve default route host names.
	IngressVIP string `json:"ingress_vip,omitempty"`

	NodeMem string `json:"node_mem,omitempty"`
	NodeCores string `json:"node_cores,omitempty"`
}
