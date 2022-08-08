// This allows sm_secret_group data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_secret_group" {
  value       = ibm_sm_secret_group.sm_secret_group_instance
  description = "sm_secret_group resource instance"
}
// This allows sm_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_secret" {
  value       = ibm_sm_secret.sm_secret_instance
  description = "sm_secret resource instance"
}
