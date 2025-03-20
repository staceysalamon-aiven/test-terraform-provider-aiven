resource "aiven_organizational_unit" "example_unit" {
  name      = "Example organizational unit - changed name for testing"
  parent_id = aiven_organization.main.org_id
}