package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"

	"github.com/hashicorp/terraform/helper/resource"
)

// Import function does not support specified provider.
func SkipTestAccAlicloudRouterInterfaceConnection_import(t *testing.T) {
	resourceName := "alicloud_router_interface_connection.foo"

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRouterInterfaceConnectionConfigBasic(acctest.RandIntRange(1000, 9999)),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
