package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIamResourceGroup_importBasic(t *testing.T) {
	resourceGroupName1 := acctest.RandomWithPrefix(test_prefix)

	resourceUrn := "urn:v1:eu:resource:vps:" + os.Getenv("OVH_VPS")

	config := fmt.Sprintf(
		`resource "ovh_iam_resource_group" "resource_group_1" {
			name        = "%s"
			resources   = ["%s"]
		}		
		`,
		resourceGroupName1,
		resourceUrn,
	)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckIamResourceGroup(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				ResourceName:      "ovh_iam_resource_group.resource_group_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
