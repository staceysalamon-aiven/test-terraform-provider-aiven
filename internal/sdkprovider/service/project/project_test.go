package project_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
)

func TestAccAivenProject_basic(t *testing.T) {
	resourceName := "aiven_project.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	rName2 := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenProjectResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenProjectAttributes("data.aiven_project.project"),
					resource.TestCheckResourceAttr(resourceName, "project", fmt.Sprintf("test-acc-pr-%s", rName)),
					resource.TestCheckResourceAttrSet(resourceName, "default_cloud"),
					resource.TestCheckResourceAttrSet(resourceName, "ca_cert"),
				),
			},
			{
				Config: testAccProjectCopyFromProjectResource(rName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenProjectAttributes("data.aiven_project.project"),
					resource.TestCheckResourceAttr(resourceName, "project", fmt.Sprintf("test-acc-pr-%s", rName2)),
					resource.TestCheckResourceAttrSet(resourceName, "default_cloud"),
					resource.TestCheckResourceAttrSet(resourceName, "ca_cert"),
					resource.TestCheckResourceAttrSet(resourceName, "billing_group"),
				),
			},
			{
				Config:             testAccProjectDoubleTagResource(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("tag keys should be unique"),
			},
		},
	})
}

func TestAccAivenProject_accounts(t *testing.T) {
	resourceName := "aiven_project.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	config := testAccProjectResourceAccounts(rName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenProjectResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenProjectAttributes("data.aiven_project.project", "account_id"),
					resource.TestCheckResourceAttr(resourceName, "project", fmt.Sprintf("test-acc-pr-%s", rName)),
				),
			},
			{
				// Tests account_id (deprecated) -> parent_id migration
				Config: regexp.MustCompile(`account_id\s+=`).ReplaceAllString(config, "parent_id ="),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "account_id"), // it is computed and should be set
					resource.TestCheckResourceAttrSet(resourceName, "parent_id"),
				),
			},
		},
	})
}

func TestAccAivenProject_organizations(t *testing.T) {
	resourceName := "aiven_project.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenProjectResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccProjectResourceOrganizations(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenProjectAttributes(resourceName, "parent_id"),
					resource.TestCheckResourceAttr(resourceName, "project", fmt.Sprintf("test-acc-pr-%s", rName)),
				),
			},
		},
	})
}

func testAccProjectDoubleTagResource(name string) string {
	return fmt.Sprintf(`
resource "aiven_account" "foo" {
  name = "test-acc-ac-%s"
}

resource "aiven_project" "foo" {
  project       = "test-acc-pr-%s"
  account_id    = aiven_account.foo.account_id
  default_cloud = "aws-eu-west-2"
  tag {
    key   = "test"
    value = "val"
  }
  tag {
    key   = "test"
    value = "val2"
  }
}

data "aiven_project" "project" {
  project = aiven_project.foo.project

  depends_on = [aiven_project.foo]
}`, name, name)
}

func testAccProjectResourceAccounts(name string) string {
	return fmt.Sprintf(`
resource "aiven_account" "foo" {
  name = "test-acc-ac-%s"
}

resource "aiven_project" "foo" {
  project       = "test-acc-pr-%s"
  account_id    = aiven_account.foo.account_id
  default_cloud = "aws-eu-west-2"
  tag {
    key   = "test"
    value = "val"
  }
}

data "aiven_project" "project" {
  project = aiven_project.foo.project

  depends_on = [aiven_project.foo]
}`, name, name)
}

func testAccProjectResourceOrganizations(name string) string {
	return fmt.Sprintf(`
resource "aiven_organization" "foo" {
  name = "test-acc-org-%s"
}

resource "aiven_project" "foo" {
  project       = "test-acc-pr-%s"
  parent_id     = aiven_organization.foo.id
  default_cloud = "aws-eu-west-2"
  tag {
    key   = "test"
    value = "val"
  }
}`, name, name)
}

func testAccProjectResource(name string) string {
	return fmt.Sprintf(`
resource "aiven_account" "bar" {
  name = "test-acc-ac-%[1]s"
}

resource "aiven_project" "foo" {
  project       = "test-acc-pr-%[1]s"
  account_id    = aiven_account.bar.account_id
  default_cloud = "aws-eu-west-2"
  tag {
    key   = "test"
    value = "val"
  }
}

data "aiven_project" "project" {
  project    = aiven_project.foo.project
  depends_on = [aiven_project.foo]
}`, name)
}

func testAccProjectCopyFromProjectResource(name string) string {
	return fmt.Sprintf(`
resource "aiven_account" "bar" {
  name = "test-acc-ac-%[1]s"
}

resource "aiven_billing_group" "foo" {
  name             = "test-acc-bg-%[1]s"
  billing_currency = "USD"
  vat_id           = "123"
}

resource "aiven_project" "source" {
  project       = "test-acc-pr-source-%[1]s"
  account_id    = aiven_account.bar.account_id
  billing_group = aiven_billing_group.foo.id
  tag {
    key   = "test"
    value = "val"
  }
}

resource "aiven_project" "foo" {
  project           = "test-acc-pr-%[1]s"
  account_id        = aiven_account.bar.account_id
  copy_from_project = aiven_project.source.project
}

data "aiven_project" "project" {
  project    = aiven_project.foo.project
  depends_on = [aiven_project.foo]
}`, name)
}

func testAccCheckAivenProjectAttributes(n string, attributes ...string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		log.Printf("[DEBUG] project attributes %v", a)

		if a["project"] == "" {
			return fmt.Errorf("expected to get a project name from Aiven")
		}

		if a["ca_cert"] == "" {
			return fmt.Errorf("expected to get an ca_cert from Aiven")
		}

		for _, attr := range attributes {
			if a[attr] == "" {
				return fmt.Errorf("expected to get an %s from Aiven", attr)
			}
		}

		return nil
	}
}

func testAccCheckAivenProjectResourceDestroy(s *terraform.State) error {
	c := acc.GetTestAivenClient()

	ctx := context.Background()

	// loop through the resources in state, verifying each project is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_project" {
			continue
		}

		p, err := c.Projects.Get(ctx, rs.Primary.ID)
		if err != nil {
			var e aiven.Error
			if errors.As(err, &e) && e.Status != 404 && e.Status != 403 {
				return err
			}
		}

		if p != nil {
			return fmt.Errorf("project (%s) still exists", rs.Primary.ID)
		}
	}

	return nil
}
