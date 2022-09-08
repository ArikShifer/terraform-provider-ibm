// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/secrets-manager-mt-go-sdk/secretsmanagerv2"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func DataSourceIbmSmSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmSecretsRead,

		Schema: map[string]*schema.Schema{
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A url pointer to the first page in a collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A url pointer to a page in a collection.",
						},
					},
				},
			},
			"previous": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A url pointer to the previous page in a collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A url pointer to a page in a collection.",
						},
					},
				},
			},
			"last": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A url pointer to the last page in a collection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A url pointer to a page in a collection.",
						},
					},
				},
			},
			"secrets": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A collection of secrets metadata.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A v4 UUID identifier.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret type. Supported types are Imported Certificate, Public Certificate.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable name to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.",
						},
						"secret_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A v4 UUID identifier.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Labels that you can use to filter for secrets in your instance.Up to 30 labels can be created.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for the entity that created the secret.",
						},
						"creation_date": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date a resource was created. The date format follows RFC 3339.",
						},
						"last_update_date": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date a resource was recently modified. The date format follows RFC 3339.",
						},
						"version_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A v4 UUID identifier.",
						},
						"versions_total": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of versions the secret has.",
						},
						"locks_total": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of locks the secret has.",
						},
						"expiration_date": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date a secret is expired. The date format follows RFC 3339.",
						},
						"intermediate_included": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the certificate was imported with an associated intermediate certificate.",
						},
						"private_key_included": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the certificate was imported with an associated private key.",
						},
						"serial_number": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique serial number that was assigned to a certificate by the issuing certificate authority.",
						},
						"algorithm": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identifier for the cryptographic algorithm that was used by the issuing certificate authority to sign a certificate.",
						},
						"key_algorithm": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The identifier for the cryptographic algorithm that was used to generate the public and private keys that are associated with the certificate.",
						},
						"issuer": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The distinguished name that identifies the entity that signed and issued the certificate.",
						},
						"validity": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The date and time that the certificate validity period begins and ends.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"not_before": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date time format follows RFC 3339.",
									},
									"not_after": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Date time format follows RFC 3339.",
									},
								},
							},
						},
						"bundle_certs": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate.",
						},
						"rotation": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Determines whether Secrets Manager rotates your secrets automatically.For public certificates, if `auto_rotate` is set to `true` the service reorders your certificate 31 daysbefore it expires.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_rotate": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Determines whether Secrets Manager rotates your public certificate automatically.Default is `false`. If `auto_rotate` is set to `true` the service reorders your certificate 31 days. If rotation fails the service will attempt to reorder your certificate on the next day, every day before expiration.",
									},
									"rotate_keys": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Determines whether Secrets Manager rotates the private key for your public certificate automatically.Default is `false`. If set to `true`, the service generates and stores a new private key for your rotated certificate.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmSecretsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	listSecretsOptions := &secretsmanagerv2.ListSecretsOptions{}

	var pager *secretsmanagerv2.SecretsPager
	pager, err = secretsManagerClient.NewSecretsPager(listSecretsOptions)
	if err != nil {
		return diag.FromErr(err)
	}

	allItems, err := pager.GetAll()
	if err != nil {
		log.Printf("[DEBUG] SecretsPager.GetAll() failed %s", err)
		return diag.FromErr(fmt.Errorf("SecretsPager.GetAll() failed %s", err))
	}

	d.SetId(dataSourceIbmSmSecretsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := dataSourceIbmSmSecretsSecretMetadataToMap(modelItem)
		if err != nil {
			return diag.FromErr(err)
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("secrets", mapSlice); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting secrets %s", err))
	}

	return nil
}

// dataSourceIbmSmSecretsID returns a reasonable ID for the list.
func dataSourceIbmSmSecretsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSmSecretsSecretMetadataPaginatedCollectionFirstToMap(model *secretsmanagerv2.SecretMetadataPaginatedCollectionFirst) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsSecretMetadataPaginatedCollectionPreviousToMap(model *secretsmanagerv2.SecretMetadataPaginatedCollectionPrevious) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsSecretMetadataPaginatedCollectionLastToMap(model *secretsmanagerv2.SecretMetadataPaginatedCollectionLast) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsSecretMetadataToMap(model secretsmanagerv2.SecretMetadataIntf) (map[string]interface{}, error) {
	if _, ok := model.(*secretsmanagerv2.ImportedCertificateMetadata); ok {
		return dataSourceIbmSmSecretsImportedCertificateMetadataToMap(model.(*secretsmanagerv2.ImportedCertificateMetadata))
	} else if _, ok := model.(*secretsmanagerv2.PublicCertificateMetadata); ok {
		return dataSourceIbmSmSecretsPublicCertificateMetadataToMap(model.(*secretsmanagerv2.PublicCertificateMetadata))
	} else if _, ok := model.(*secretsmanagerv2.SecretMetadata); ok {
		modelMap := make(map[string]interface{})
		model := model.(*secretsmanagerv2.SecretMetadata)
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.Type != nil {
			modelMap["type"] = *model.Type
		}
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.Description != nil {
			modelMap["description"] = *model.Description
		}
		if model.SecretGroupID != nil {
			modelMap["secret_group_id"] = *model.SecretGroupID
		}
		if model.Labels != nil {
			modelMap["labels"] = model.Labels
		}
		if model.CreatedBy != nil {
			modelMap["created_by"] = *model.CreatedBy
		}
		if model.CreationDate != nil {
			modelMap["creation_date"] = model.CreationDate.String()
		}
		if model.LastUpdateDate != nil {
			modelMap["last_update_date"] = model.LastUpdateDate.String()
		}
		if model.VersionID != nil {
			modelMap["version_id"] = *model.VersionID
		}
		if model.VersionsTotal != nil {
			modelMap["versions_total"] = *model.VersionsTotal
		}
		if model.LocksTotal != nil {
			modelMap["locks_total"] = *model.LocksTotal
		}
		if model.ExpirationDate != nil {
			modelMap["expiration_date"] = model.ExpirationDate.String()
		}
		if model.IntermediateIncluded != nil {
			modelMap["intermediate_included"] = *model.IntermediateIncluded
		}
		if model.PrivateKeyIncluded != nil {
			modelMap["private_key_included"] = *model.PrivateKeyIncluded
		}
		if model.SerialNumber != nil {
			modelMap["serial_number"] = *model.SerialNumber
		}
		if model.Algorithm != nil {
			modelMap["algorithm"] = *model.Algorithm
		}
		if model.KeyAlgorithm != nil {
			modelMap["key_algorithm"] = *model.KeyAlgorithm
		}
		if model.Issuer != nil {
			modelMap["issuer"] = *model.Issuer
		}
		if model.Validity != nil {
			validityMap, err := dataSourceIbmSmSecretsCertificateValidityToMap(model.Validity)
			if err != nil {
				return modelMap, err
			}
			modelMap["validity"] = []map[string]interface{}{validityMap}
		}
		if model.BundleCerts != nil {
			modelMap["bundle_certs"] = *model.BundleCerts
		}
		if model.Rotation != nil {
			rotationMap, err := dataSourceIbmSmSecretsPublicCertificateRotationPolicyToMap(model.Rotation)
			if err != nil {
				return modelMap, err
			}
			modelMap["rotation"] = []map[string]interface{}{rotationMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized secretsmanagerv2.SecretMetadataIntf subtype encountered")
	}
}

func dataSourceIbmSmSecretsCertificateValidityToMap(model *secretsmanagerv2.CertificateValidity) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NotBefore != nil {
		modelMap["not_before"] = model.NotBefore.String()
	}
	if model.NotAfter != nil {
		modelMap["not_after"] = model.NotAfter.String()
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsPublicCertificateRotationPolicyToMap(model *secretsmanagerv2.PublicCertificateRotationPolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AutoRotate != nil {
		modelMap["auto_rotate"] = *model.AutoRotate
	}
	if model.RotateKeys != nil {
		modelMap["rotate_keys"] = *model.RotateKeys
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsPublicCertificateMetadataToMap(model *secretsmanagerv2.PublicCertificateMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.SecretGroupID != nil {
		modelMap["secret_group_id"] = *model.SecretGroupID
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreationDate != nil {
		modelMap["creation_date"] = model.CreationDate.String()
	}
	if model.LastUpdateDate != nil {
		modelMap["last_update_date"] = model.LastUpdateDate.String()
	}
	if model.VersionID != nil {
		modelMap["version_id"] = *model.VersionID
	}
	if model.VersionsTotal != nil {
		modelMap["versions_total"] = *model.VersionsTotal
	}
	if model.LocksTotal != nil {
		modelMap["locks_total"] = *model.LocksTotal
	}
	if model.ExpirationDate != nil {
		modelMap["expiration_date"] = model.ExpirationDate.String()
	}
	if model.SerialNumber != nil {
		modelMap["serial_number"] = *model.SerialNumber
	}
	if model.Algorithm != nil {
		modelMap["algorithm"] = *model.Algorithm
	}
	if model.KeyAlgorithm != nil {
		modelMap["key_algorithm"] = *model.KeyAlgorithm
	}
	if model.Issuer != nil {
		modelMap["issuer"] = *model.Issuer
	}
	if model.Validity != nil {
		validityMap, err := dataSourceIbmSmSecretsCertificateValidityToMap(model.Validity)
		if err != nil {
			return modelMap, err
		}
		modelMap["validity"] = []map[string]interface{}{validityMap}
	}
	if model.BundleCerts != nil {
		modelMap["bundle_certs"] = *model.BundleCerts
	}
	if model.Rotation != nil {
		rotationMap, err := dataSourceIbmSmSecretsPublicCertificateRotationPolicyToMap(model.Rotation)
		if err != nil {
			return modelMap, err
		}
		modelMap["rotation"] = []map[string]interface{}{rotationMap}
	}
	return modelMap, nil
}

func dataSourceIbmSmSecretsImportedCertificateMetadataToMap(model *secretsmanagerv2.ImportedCertificateMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.SecretGroupID != nil {
		modelMap["secret_group_id"] = *model.SecretGroupID
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CreationDate != nil {
		modelMap["creation_date"] = model.CreationDate.String()
	}
	if model.LastUpdateDate != nil {
		modelMap["last_update_date"] = model.LastUpdateDate.String()
	}
	if model.VersionID != nil {
		modelMap["version_id"] = *model.VersionID
	}
	if model.VersionsTotal != nil {
		modelMap["versions_total"] = *model.VersionsTotal
	}
	if model.LocksTotal != nil {
		modelMap["locks_total"] = *model.LocksTotal
	}
	if model.ExpirationDate != nil {
		modelMap["expiration_date"] = model.ExpirationDate.String()
	}
	if model.IntermediateIncluded != nil {
		modelMap["intermediate_included"] = *model.IntermediateIncluded
	}
	if model.PrivateKeyIncluded != nil {
		modelMap["private_key_included"] = *model.PrivateKeyIncluded
	}
	if model.SerialNumber != nil {
		modelMap["serial_number"] = *model.SerialNumber
	}
	if model.Algorithm != nil {
		modelMap["algorithm"] = *model.Algorithm
	}
	if model.KeyAlgorithm != nil {
		modelMap["key_algorithm"] = *model.KeyAlgorithm
	}
	if model.Issuer != nil {
		modelMap["issuer"] = *model.Issuer
	}
	if model.Validity != nil {
		validityMap, err := dataSourceIbmSmSecretsCertificateValidityToMap(model.Validity)
		if err != nil {
			return modelMap, err
		}
		modelMap["validity"] = []map[string]interface{}{validityMap}
	}
	return modelMap, nil
}
