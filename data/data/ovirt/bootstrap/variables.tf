variable "ignition_bootstrap" {
  description = "The ID of storage domain"
  default     = ""
}

variable "ovirt_cluster_id" {
  type        = string
  default     = ""
  description = "The name of cluster"
}

variable "cluster_domain" {
  description = "The ID of Openshift cluster"
}

variable "ovirt_template_id" {
  type        = string
  default     = ""
  description = "The ID of cluster"
}
