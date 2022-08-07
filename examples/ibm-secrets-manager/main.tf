provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision sm_secret_group resource instance
resource "ibm_sm_secret_group" "sm_secret_group_instance" {
  name = var.sm_secret_group_name
  description = var.sm_secret_group_description
}
