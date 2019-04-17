variable "bootstrap_dns" {
  type        = string
  default     = true
  description = "Whether to include DNS entries for the bootstrap node or not."
}

variable "ovirt_url" {
  type        = string
  default     = ""
  description = "The oVirt engine URL"
}

variable "ovirt_username" {
  type        = string
  default     = ""
  description = "The name of user to access oVirt engine API"
}

variable "ovirt_password" {
  type        = string
  default     = ""
  description = "The plain password of user to access oVirt engine API"
}

variable "ovirt_cafile" {
  type        = string
  default     = ""
  description = "The CA file of oVirt engine certificates"
}

variable "ovirt_cluster_id" {
  type        = string
  default     = ""
  description = "The name of cluster"
}

variable "ovirt_template_name" {
  type        = string
  default     = "rhcos"
  description = "The ID of VM template"
}

variable "ovirt_template_id" {
  type        = string
  default     = ""
  description = "The ID of cluster"
}
