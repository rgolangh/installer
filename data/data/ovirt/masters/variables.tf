variable "cluster_id" {
  description = "The ID of Openshift cluster"
}

variable "cluster_domain" {
  description = "The domain name of Openshift cluster"
}

variable "master_count" {
  type        = string
  description = "Number of masters"
  default     = 3
}

variable "ovirt_cluster_id" {
  type        = string
  description = "The ID of oVirt's cluster"
}

variable "ovirt_template_id" {
  type        = string
  description = "The ID of oVirt's VM template"
}

variable "ovirt_master_instance_type_id" {
  type        = string
  description = "The ID of oVirt's instance type"
}

variable "ovirt_master_vm_type" {
  type        = string
  description = "The master's VM type"
}

variable "ignition_master" {
  type        = string
  description = "master ignition config"
}

variable "ovirt_master_memory" {
  type = string
}

variable "ovirt_master_cores" {
  type = string
}

variable "ovirt_master_sockets" {
  type = string
}

variable "ovirt_master_os_disk_size_gb" {
  type = string
}