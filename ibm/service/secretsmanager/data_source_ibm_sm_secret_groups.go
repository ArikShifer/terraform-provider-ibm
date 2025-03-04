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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func DataSourceIbmSmSecretGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmSecretGroupsRead,

		Schema: map[string]*schema.Schema{
			"secret_groups": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A collection of secret groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A v4 UUID identifier.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of your secret group.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.",
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
					},
				},
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources in a collection.",
			},
		},
	}
}

func dataSourceIbmSmSecretGroupsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	listSecretGroupsOptions := &secretsmanagerv2.ListSecretGroupsOptions{}

	secretGroupCollection, response, err := secretsManagerClient.ListSecretGroupsWithContext(context, listSecretGroupsOptions)
	if err != nil {
		log.Printf("[DEBUG] ListSecretGroupsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListSecretGroupsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmSmSecretGroupsID(d))

	secretGroups := []map[string]interface{}{}
	if secretGroupCollection.SecretGroups != nil {
		for _, modelItem := range secretGroupCollection.SecretGroups {
			modelMap, err := dataSourceIbmSmSecretGroupsSecretGroupToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			secretGroups = append(secretGroups, modelMap)
		}
	}
	if err = d.Set("secret_groups", secretGroups); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting secret_groups %s", err))
	}

	if err = d.Set("total_count", flex.IntValue(secretGroupCollection.TotalCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting total_count: %s", err))
	}

	return nil
}

// dataSourceIbmSmSecretGroupsID returns a reasonable ID for the list.
func dataSourceIbmSmSecretGroupsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSmSecretGroupsSecretGroupToMap(model *secretsmanagerv2.SecretGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.CreationDate != nil {
		modelMap["creation_date"] = model.CreationDate.String()
	}
	if model.LastUpdateDate != nil {
		modelMap["last_update_date"] = model.LastUpdateDate.String()
	}
	return modelMap, nil
}
