// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/secrets-manager-mt-go-sdk/secretsmanagerv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmSmSecretBasic(t *testing.T) {
	var conf secretsmanagerv1.Secret

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmSecretExists("ibm_sm_secret.sm_secret", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_sm_secret.sm_secret",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSmSecretConfigBasic() string {
	return fmt.Sprintf(`

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
	`)
}

func testAccCheckIbmSmSecretExists(n string, obj secretsmanagerv1.Secret) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV1()
		if err != nil {
			return err
		}

		getSecretOptions := &secretsmanagerv1.GetSecretOptions{}

		getSecretOptions.SetID(rs.Primary.ID)

		secretIntf, _, err := secretsManagerClient.GetSecret(getSecretOptions)
		if err != nil {
			return err
		}

		secret := secretIntf.(*secretsmanagerv1.Secret)
		obj = *secret
		return nil
	}
}

func testAccCheckIbmSmSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv1.GetSecretOptions{}

		getSecretOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("sm_secret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for sm_secret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
