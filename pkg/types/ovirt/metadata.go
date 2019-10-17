package ovirt

// Metadata contains ovirt metadata (e.g. for uninstalling the cluster).
type Metadata struct {
	URL         string `json:"url"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Cafile      string `json:"cafile"`
	APIVIP      string `json:"api_vip"`
	DNSVIP      string `json:"dns_vip"`
	IngressVIP  string `json:"ingress_vip"`
	NodeMem     string `json:"node_mem"`
	NodeCores   string `json:"node_cores"`
}
