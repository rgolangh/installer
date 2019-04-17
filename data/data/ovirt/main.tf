provider "ovirt" {
  url      = var.ovirt_url
  username = var.ovirt_username
  password = var.ovirt_password
}

module "bootstrap" {
  source             = "./bootstrap"
  ovirt_cluster_id   = var.ovirt_cluster_id
  ovirt_template_id  = var.ovirt_template_id
  ignition_bootstrap = var.ignition_bootstrap
  cluster_domain     = var.cluster_domain
}

resource "ovirt_vm" "master" {
  count       = var.master_count
  name        = "master-${count.index}.${var.cluster_domain}"
  cluster_id  = var.ovirt_cluster_id
  template_id = var.ovirt_template_id
  memory      = "8192"
  cores       = "4"

  initialization {
    host_name     = "master-${count.index}.local"
    custom_script = var.ignition_master
  }
}

resource "ovirt_tag" "cluster_tag" {
  name   = var.cluster_id
  vm_ids = [for instance in ovirt_vm.master.* : instance.id]
}
