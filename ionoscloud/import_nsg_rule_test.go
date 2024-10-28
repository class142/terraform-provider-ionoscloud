package ionoscloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/utils/constant"
)

func TestAccNSGRuleImportBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ExternalProviders: randomProviderVersion343(),
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckNSGRuleDestroyCheck,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetworkSecurityGroupFirewallRulesBasic,
			},
			{
				ResourceName:      constant.NetworkSecurityGroupFirewallRuleResource + "." + constant.NetworkSecurityGroupFirewallRuleTestResource + "_1",
				ImportStateIdFunc: testAccNSGRuleImportStateId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      constant.NetworkSecurityGroupFirewallRuleResource + "." + constant.NetworkSecurityGroupFirewallRuleTestResource + "_2",
				ImportStateIdFunc: testAccNSGRuleImportStateId,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNSGRuleImportStateId(s *terraform.State) (string, error) {
	importID := ""

	for _, rs := range s.RootModule().Resources {
		if rs.Type != constant.NetworkSecurityGroupFirewallRuleResource {
			continue
		}

		importID = fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes["datacenter_id"], rs.Primary.Attributes["nsg_id"], rs.Primary.Attributes["id"])
	}

	return importID, nil
}