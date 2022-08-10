# Example for SecretsManagerV2

This example illustrates how to use the SecretsManagerV2

These types of resources are supported:

* SecretGroup
* sm_secret

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## SecretsManagerV2 resources

sm_secret_group resource:

```hcl
resource "sm_secret_group" "sm_secret_group_instance" {
  name = var.sm_secret_group_name
  description = var.sm_secret_group_description
}
```
sm_secret resource:

```hcl
resource "sm_secret" "sm_secret_instance" {
  secret_prototype = var.sm_secret_secret_prototype
}
```

## SecretsManagerV2 Data sources


## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of your secret group. | `string` | true |
| description | An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group. | `string` | false |
| secret_prototype | Specify the properties for your secret. | `` | true |

## Outputs

| Name | Description |
|------|-------------|
| sm_secret_group | sm_secret_group object |
| sm_secret | sm_secret object |
