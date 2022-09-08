---
layout: "ibm"
page_title: "IBM : ibm_sm_secret"
description: |-
  Manages sm_secret.
subcategory: "IBM Cloud Secrets Manager Basic API"
---

# ibm_sm_secret

Provides a resource for sm_secret. This allows sm_secret to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_sm_secret" "sm_secret" {
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
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `secret_prototype` - (Required, Forces new resource, List) Specify the properties for your secret.
Nested scheme for **secret_prototype**:
	* `bundle_certs` - (Optional, Boolean) Determines whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate.
	  * Constraints: The default value is `true`.
	* `certificate` - (Optional, String) The PEM encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
	* `description` - (Optional, String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
	* `intermediate` - (Optional, String) (Optional) The PEM encoded intermediate certificate to associate with the root certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
	* `labels` - (Optional, List) Labels that you can use to filter for secrets in your instance.Up to 30 labels can be created.
	  * Constraints: The list items must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`. The maximum length is `30` items. The minimum length is `0` items.
	* `name` - (Optional, String) A human-readable name to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.
	  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.
	* `private_key` - (Optional, String) (Optional) The PEM encoded private key to associate with the certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
	* `rotation` - (Optional, List) Determines whether Secrets Manager rotates your secrets automatically.For public certificates, if `auto_rotate` is set to `true` the service reorders your certificate 31 daysbefore it expires.
	Nested scheme for **rotation**:
		* `auto_rotate` - (Optional, Boolean) Determines whether Secrets Manager rotates your public certificate automatically.Default is `false`. If `auto_rotate` is set to `true` the service reorders your certificate 31 days. If rotation fails the service will attempt to reorder your certificate on the next day, every day before expiration.
		  * Constraints: The default value is `false`.
		* `rotate_keys` - (Optional, Boolean) Determines whether Secrets Manager rotates the private key for your public certificate automatically.Default is `false`. If set to `true`, the service generates and stores a new private key for your rotated certificate.
		  * Constraints: The default value is `false`.
	* `secret_group_id` - (Optional, String) A v4 UUID identifier.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
	* `type` - (Optional, String) Secret type. Supported types are Imported Certificate, Public Certificate.
	  * Constraints: Allowable values are: `imported_cert`, `public_cert`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the sm_secret.
* `algorithm` - (String) The identifier for the cryptographic algorithm that was used by the issuing certificate authority to sign a certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters.
* `bundle_certs` - (Boolean) Determines whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate.
  * Constraints: The default value is `true`.
* `created_by` - (String) The unique identifier for the entity that created the secret.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `creation_date` - (String) The date a resource was created. The date format follows RFC 3339.
* `description` - (String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
* `intermediate_included` - (Boolean) Indicates whether the certificate was imported with an associated intermediate certificate.
* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters.
* `key_algorithm` - (String) The identifier for the cryptographic algorithm that was used to generate the public and private keys that are associated with the certificate.
  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters.
* `labels` - (List) Labels that you can use to filter for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`. The maximum length is `30` items. The minimum length is `0` items.
* `last_update_date` - (String) The date a resource was recently modified. The date format follows RFC 3339.
* `locks_total` - (Integer) The number of locks the secret has.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.
* `name` - (String) A human-readable name to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.
* `private_key_included` - (Boolean) Indicates whether the certificate was imported with an associated private key.
* `rotation` - (List) Determines whether Secrets Manager rotates your secrets automatically.For public certificates, if `auto_rotate` is set to `true` the service reorders your certificate 31 daysbefore it expires.
Nested scheme for **rotation**:
	* `auto_rotate` - (Boolean) Determines whether Secrets Manager rotates your public certificate automatically.Default is `false`. If `auto_rotate` is set to `true` the service reorders your certificate 31 days. If rotation fails the service will attempt to reorder your certificate on the next day, every day before expiration.
	  * Constraints: The default value is `false`.
	* `rotate_keys` - (Boolean) Determines whether Secrets Manager rotates the private key for your public certificate automatically.Default is `false`. If set to `true`, the service generates and stores a new private key for your rotated certificate.
	  * Constraints: The default value is `false`.
* `secret_data` - (List) Your secret data.
Nested scheme for **secret_data**:
	* `certificate` - (String) The PEM encoded contents of your certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
	* `intermediate` - (String) (Optional) The PEM encoded intermediate certificate to associate with the root certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
	* `private_key` - (String) (Optional) The PEM encoded private key to associate with the certificate.
	  * Constraints: The maximum length is `100000` characters. The minimum length is `50` characters. The value must match regular expression `/((.| )*)/`.
* `secret_group_id` - (String) A v4 UUID identifier.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `serial_number` - (String) The unique serial number that was assigned to a certificate by the issuing certificate authority.
  * Constraints: The maximum length is `64` characters. The minimum length is `32` characters.
* `type` - (String) Secret type. Supported types are Imported Certificate, Public Certificate.
  * Constraints: Allowable values are: `imported_cert`, `public_cert`.
* `validity` - (List) The date and time that the certificate validity period begins and ends.
Nested scheme for **validity**:
	* `not_after` - (String) Date time format follows RFC 3339.
	* `not_before` - (String) Date time format follows RFC 3339.
* `version_id` - (String) A v4 UUID identifier.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
* `versions_total` - (Integer) The number of versions the secret has.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_secret` resource by using `id`. A v4 UUID identifier.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```
$ terraform import ibm_sm_secret.sm_secret <id>
```

# Example
```
$ terraform import ibm_sm_secret.sm_secret b49ad24d-81d4-5ebc-b9b9-b0937d1c84d5
```
