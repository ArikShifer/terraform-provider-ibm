// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmSecretsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmSecretsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_secrets.sm_secrets", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secrets.sm_secrets", "secrets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_secrets.sm_secrets", "total_count"),
				),
			},
		},
	})
}

func testAccCheckIbmSmSecretsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_sm_secrets" "sm_secrets" {
		}
	`)
}
