provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision sm_secret_group resource instance
resource "ibm_sm_secret_group" "sm_secret_group_instance" {
  name = var.sm_secret_group_name
  description = var.sm_secret_group_description
}

// Provision sm_secret resource instance
resource "ibm_sm_secret" "sm_secret_instance" {
  secret_prototype {
    type = "imported_cert"
    name = "my-secret"
    description = "Extended description for this secret."
    secret_group_id = "b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5"
    labels = [ "my-label" ]
    certificate = "certificate"
    intermediate = "intermediate"
    private_key = "private_key"
  }
}
