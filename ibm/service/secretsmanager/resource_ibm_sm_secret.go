// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/secrets-manager-mt-go-sdk/secretsmanagerv1"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmSmSecret() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmSecretCreate,
		ReadContext:   resourceIbmSmSecretRead,
		DeleteContext: resourceIbmSmSecretDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"secret_prototype": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				ForceNew:    true,
				Description: "Specify the properties for your secret.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secret type. Supported types are Imported Certificate, Public Certificate.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A human-readable name to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as an name for your secret.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.",
						},
						"secret_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A v4 UUID identifier.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Labels that you can use to filter for secrets in your instance.Up to 30 labels can be created.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"certificate": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The PEM encoded contents of your certificate.",
						},
						"intermediate": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "(Optional) The PEM encoded intermediate certificate to associate with the root certificate.",
						},
						"private_key": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "(Optional) The PEM encoded private key to associate with the certificate.",
						},
						"bundle_certs": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Determines whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate.",
						},
					},
				},
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
				Elem:        &schema.Schema{Type: schema.TypeString},
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
			"versions_total": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of versions the secret has.",
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
							Required:    true,
							Description: "Date time format follows RFC 3339.",
						},
						"not_after": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Date time format follows RFC 3339.",
						},
					},
				},
			},
			"certificate": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The PEM encoded contents of your certificate.",
			},
			"intermediate": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "(Optional) The PEM encoded intermediate certificate to associate with the root certificate.",
			},
			"private_key": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "(Optional) The PEM encoded private key to associate with the certificate.",
			},
			"bundle_certs": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines whether your issued certificate is bundled with intermediate certificates. Set to `false` for the certificate file to contain only the issued certificate.",
			},
		},
	}
}

func resourceIbmSmSecretCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createSecretOptions := &secretsmanagerv1.CreateSecretOptions{}

	secretPrototypeModel, err := resourceIbmSmSecretMapToSecretPrototype(d.Get("secret_prototype.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createSecretOptions.SetSecretPrototype(secretPrototypeModel)

	secretIntf, response, err := secretsManagerClient.CreateSecretWithContext(context, createSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateSecretWithContext failed %s\n%s", err, response))
	}

	secret := secretIntf.(*secretsmanagerv1.ImportedCertificate)
	d.SetId(*secret.ID)

	return resourceIbmSmSecretRead(context, d, meta)
}

func resourceIbmSmSecretRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getSecretOptions := &secretsmanagerv1.GetSecretOptions{}

	getSecretOptions.SetID(d.Id())

	secretIntf, response, err := secretsManagerClient.GetSecretWithContext(context, getSecretOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSecretWithContext failed %s\n%s", err, response))
	}

	if _, ok := secretIntf.(*secretsmanagerv1.ImportedCertificate); ok {
		secret := secretIntf.(*secretsmanagerv1.ImportedCertificate)
		if err = d.Set("type", secret.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
		if err = d.Set("name", secret.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
		if err = d.Set("description", secret.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
		if err = d.Set("secret_group_id", secret.SecretGroupID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret_group_id: %s", err))
		}
		if secret.Labels != nil {
			if err = d.Set("labels", secret.Labels); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting labels: %s", err))
			}
		}
		if err = d.Set("created_by", secret.CreatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
		}
		if err = d.Set("creation_date", flex.DateTimeToString(secret.CreationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting creation_date: %s", err))
		}
		if err = d.Set("last_update_date", flex.DateTimeToString(secret.LastUpdateDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_update_date: %s", err))
		}
		if err = d.Set("versions_total", flex.IntValue(secret.VersionsTotal)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting versions_total: %s", err))
		}
		if err = d.Set("expiration_date", flex.DateTimeToString(secret.ExpirationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting expiration_date: %s", err))
		}
		if err = d.Set("intermediate_included", secret.IntermediateIncluded); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting intermediate_included: %s", err))
		}
		if err = d.Set("private_key_included", secret.PrivateKeyIncluded); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting private_key_included: %s", err))
		}
		if err = d.Set("serial_number", secret.SerialNumber); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting serial_number: %s", err))
		}
		if err = d.Set("algorithm", secret.Algorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting algorithm: %s", err))
		}
		if err = d.Set("key_algorithm", secret.KeyAlgorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting key_algorithm: %s", err))
		}
		if err = d.Set("issuer", secret.Issuer); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting issuer: %s", err))
		}
		if secret.Validity != nil {
			validityMap, err := resourceIbmSmSecretCertificateValidityToMap(secret.Validity)
			if err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("validity", []map[string]interface{}{validityMap}); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting validity: %s", err))
			}
		}
		if err = d.Set("certificate", secret.Certificate); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting certificate: %s", err))
		}
		if err = d.Set("intermediate", secret.Intermediate); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting intermediate: %s", err))
		}
		if err = d.Set("private_key", secret.PrivateKey); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting private_key: %s", err))
		}
	} else if _, ok := secretIntf.(*secretsmanagerv1.PublicCertificate); ok {
		secret := secretIntf.(*secretsmanagerv1.PublicCertificate)
		if err = d.Set("type", secret.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
		if err = d.Set("name", secret.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
		if err = d.Set("description", secret.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
		if err = d.Set("secret_group_id", secret.SecretGroupID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret_group_id: %s", err))
		}
		if secret.Labels != nil {
			if err = d.Set("labels", secret.Labels); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting labels: %s", err))
			}
		}
		if err = d.Set("created_by", secret.CreatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
		}
		if err = d.Set("creation_date", flex.DateTimeToString(secret.CreationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting creation_date: %s", err))
		}
		if err = d.Set("last_update_date", flex.DateTimeToString(secret.LastUpdateDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_update_date: %s", err))
		}
		if err = d.Set("versions_total", flex.IntValue(secret.VersionsTotal)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting versions_total: %s", err))
		}
		if err = d.Set("expiration_date", flex.DateTimeToString(secret.ExpirationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting expiration_date: %s", err))
		}
		if err = d.Set("serial_number", secret.SerialNumber); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting serial_number: %s", err))
		}
		if err = d.Set("algorithm", secret.Algorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting algorithm: %s", err))
		}
		if err = d.Set("key_algorithm", secret.KeyAlgorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting key_algorithm: %s", err))
		}
		if err = d.Set("issuer", secret.Issuer); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting issuer: %s", err))
		}
		if secret.Validity != nil {
			validityMap, err := resourceIbmSmSecretCertificateValidityToMap(secret.Validity)
			if err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("validity", []map[string]interface{}{validityMap}); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting validity: %s", err))
			}
		}
		if err = d.Set("bundle_certs", secret.BundleCerts); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting bundle_certs: %s", err))
		}
	} else if _, ok := secretIntf.(*secretsmanagerv1.Secret); ok {
		secret := secretIntf.(*secretsmanagerv1.Secret)
		// TODO: handle argument of type SecretPrototype
		if err = d.Set("type", secret.Type); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
		}
		if err = d.Set("name", secret.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
		if err = d.Set("description", secret.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
		if err = d.Set("secret_group_id", secret.SecretGroupID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting secret_group_id: %s", err))
		}
		if secret.Labels != nil {
			if err = d.Set("labels", secret.Labels); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting labels: %s", err))
			}
		}
		if err = d.Set("created_by", secret.CreatedBy); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_by: %s", err))
		}
		if err = d.Set("creation_date", flex.DateTimeToString(secret.CreationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting creation_date: %s", err))
		}
		if err = d.Set("last_update_date", flex.DateTimeToString(secret.LastUpdateDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_update_date: %s", err))
		}
		if err = d.Set("versions_total", flex.IntValue(secret.VersionsTotal)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting versions_total: %s", err))
		}
		if err = d.Set("expiration_date", flex.DateTimeToString(secret.ExpirationDate)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting expiration_date: %s", err))
		}
		if err = d.Set("intermediate_included", secret.IntermediateIncluded); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting intermediate_included: %s", err))
		}
		if err = d.Set("private_key_included", secret.PrivateKeyIncluded); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting private_key_included: %s", err))
		}
		if err = d.Set("serial_number", secret.SerialNumber); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting serial_number: %s", err))
		}
		if err = d.Set("algorithm", secret.Algorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting algorithm: %s", err))
		}
		if err = d.Set("key_algorithm", secret.KeyAlgorithm); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting key_algorithm: %s", err))
		}
		if err = d.Set("issuer", secret.Issuer); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting issuer: %s", err))
		}
		if secret.Validity != nil {
			validityMap, err := resourceIbmSmSecretCertificateValidityToMap(secret.Validity)
			if err != nil {
				return diag.FromErr(err)
			}
			if err = d.Set("validity", []map[string]interface{}{validityMap}); err != nil {
				return diag.FromErr(fmt.Errorf("Error setting validity: %s", err))
			}
		}
		if err = d.Set("certificate", secret.Certificate); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting certificate: %s", err))
		}
		if err = d.Set("intermediate", secret.Intermediate); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting intermediate: %s", err))
		}
		if err = d.Set("private_key", secret.PrivateKey); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting private_key: %s", err))
		}
		if err = d.Set("bundle_certs", secret.BundleCerts); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting bundle_certs: %s", err))
		}
	} else {
		return diag.FromErr(fmt.Errorf("Unrecognized secretsmanagerv1.SecretIntf subtype encountered"))
	}

	return nil
}

func resourceIbmSmSecretDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteSecretOptions := &secretsmanagerv1.DeleteSecretOptions{}

	deleteSecretOptions.SetID(d.Id())

	response, err := secretsManagerClient.DeleteSecretWithContext(context, deleteSecretOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteSecretWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteSecretWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmSmSecretMapToSecretPrototype(modelMap map[string]interface{}) (secretsmanagerv1.SecretPrototypeIntf, error) {
	discValue, ok := modelMap["type"]
	if ok {
		if discValue == "imported_cert" {
			return resourceIbmSmSecretMapToImportedCertificatePrototype(modelMap)
		} else if discValue == "public_cert" {
			return resourceIbmSmSecretMapToPublicCertificatePrototype(modelMap)
		} else {
			return nil, fmt.Errorf("unexpected value for discriminator property 'type' found in map: '%s'", discValue)
		}
	} else {
		return nil, fmt.Errorf("discriminator property 'type' not found in map")
	}
}

func resourceIbmSmSecretMapToPublicCertificatePrototype(modelMap map[string]interface{}) (*secretsmanagerv1.PublicCertificatePrototype, error) {
	model := &secretsmanagerv1.PublicCertificatePrototype{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["secret_group_id"] != nil && modelMap["secret_group_id"].(string) != "" {
		model.SecretGroupID = core.StringPtr(modelMap["secret_group_id"].(string))
	}
	if modelMap["labels"] != nil {
		labels := []string{}
		for _, labelsItem := range modelMap["labels"].([]interface{}) {
			labels = append(labels, labelsItem.(string))
		}
		model.Labels = labels
	}
	model.BundleCerts = core.BoolPtr(modelMap["bundle_certs"].(bool))
	return model, nil
}

func resourceIbmSmSecretMapToImportedCertificatePrototype(modelMap map[string]interface{}) (*secretsmanagerv1.ImportedCertificatePrototype, error) {
	model := &secretsmanagerv1.ImportedCertificatePrototype{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["secret_group_id"] != nil && modelMap["secret_group_id"].(string) != "" {
		model.SecretGroupID = core.StringPtr(modelMap["secret_group_id"].(string))
	}
	if modelMap["labels"] != nil {
		labels := []string{}
		for _, labelsItem := range modelMap["labels"].([]interface{}) {
			labels = append(labels, labelsItem.(string))
		}
		model.Labels = labels
	}
	model.Certificate = core.StringPtr(modelMap["certificate"].(string))
	if modelMap["intermediate"] != nil && modelMap["intermediate"].(string) != "" {
		model.Intermediate = core.StringPtr(modelMap["intermediate"].(string))
	}
	if modelMap["private_key"] != nil && modelMap["private_key"].(string) != "" {
		model.PrivateKey = core.StringPtr(modelMap["private_key"].(string))
	}
	return model, nil
}

func resourceIbmSmSecretSecretPrototypeToMap(model secretsmanagerv1.SecretPrototypeIntf) (map[string]interface{}, error) {
	if _, ok := model.(*secretsmanagerv1.ImportedCertificatePrototype); ok {
		return resourceIbmSmSecretImportedCertificatePrototypeToMap(model.(*secretsmanagerv1.ImportedCertificatePrototype))
	} else if _, ok := model.(*secretsmanagerv1.PublicCertificatePrototype); ok {
		return resourceIbmSmSecretPublicCertificatePrototypeToMap(model.(*secretsmanagerv1.PublicCertificatePrototype))
	} else if _, ok := model.(*secretsmanagerv1.SecretPrototype); ok {
		modelMap := make(map[string]interface{})
		model := model.(*secretsmanagerv1.SecretPrototype)
		if model.Type != nil {
			modelMap["type"] = model.Type
		}
		if model.Name != nil {
			modelMap["name"] = model.Name
		}
		if model.Description != nil {
			modelMap["description"] = model.Description
		}
		if model.SecretGroupID != nil {
			modelMap["secret_group_id"] = model.SecretGroupID
		}
		if model.Labels != nil {
			modelMap["labels"] = model.Labels
		}
		if model.Certificate != nil {
			modelMap["certificate"] = model.Certificate
		}
		if model.Intermediate != nil {
			modelMap["intermediate"] = model.Intermediate
		}
		if model.PrivateKey != nil {
			modelMap["private_key"] = model.PrivateKey
		}
		if model.BundleCerts != nil {
			modelMap["bundle_certs"] = model.BundleCerts
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized secretsmanagerv1.SecretPrototypeIntf subtype encountered")
	}
}

func resourceIbmSmSecretPublicCertificatePrototypeToMap(model *secretsmanagerv1.PublicCertificatePrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.SecretGroupID != nil {
		modelMap["secret_group_id"] = model.SecretGroupID
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	modelMap["bundle_certs"] = model.BundleCerts
	return modelMap, nil
}

func resourceIbmSmSecretImportedCertificatePrototypeToMap(model *secretsmanagerv1.ImportedCertificatePrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.SecretGroupID != nil {
		modelMap["secret_group_id"] = model.SecretGroupID
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	modelMap["certificate"] = model.Certificate
	if model.Intermediate != nil {
		modelMap["intermediate"] = model.Intermediate
	}
	if model.PrivateKey != nil {
		modelMap["private_key"] = model.PrivateKey
	}
	return modelMap, nil
}

func resourceIbmSmSecretCertificateValidityToMap(model *secretsmanagerv1.CertificateValidity) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["not_before"] = model.NotBefore.String()
	modelMap["not_after"] = model.NotAfter.String()
	return modelMap, nil
}
