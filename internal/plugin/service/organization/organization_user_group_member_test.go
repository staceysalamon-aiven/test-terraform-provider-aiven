package organization_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/aiven/terraform-provider-aiven/internal/plugin/util"
)

// TestAccOrganizationUserGroupMember tests the organization user group member resource.
func TestAccOrganizationUserGroupMember(t *testing.T) {
	deps := acc.CommonTestDependencies(t)

	_ = deps.IsBeta(true)

	userID := deps.OrganizationUserID(true)

	name := "aiven_organization_user_group_member.foo"

	suffix := acctest.RandStringFromCharSet(acc.DefaultRandomSuffixLength, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "aiven_organization" "foo" {
  name = "%[3]s"
}

resource "aiven_organization_user_group" "foo" {
  organization_id = data.aiven_organization.foo.id
  name            = "%[1]s-usr-group-%[2]s"
  description     = "Terraform acceptance tests"
}

resource "aiven_organization_user_group_member" "foo" {
  organization_id = data.aiven_organization.foo.id
  group_id        = aiven_organization_user_group.foo.group_id
  user_id         = "%[4]s"
}
	`, acc.DefaultResourceNamePrefix, suffix, deps.OrganizationName(), *userID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "last_activity_time"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					rs, err := acc.ResourceFromState(state, name)
					if err != nil {
						return "", err
					}

					return util.ComposeID(
						rs.Primary.Attributes["organization_id"],
						rs.Primary.Attributes["group_id"],
						rs.Primary.Attributes["user_id"],
					), nil
				},
			},
		},
	})
}
