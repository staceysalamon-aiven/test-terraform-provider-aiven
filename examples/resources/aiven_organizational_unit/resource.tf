resource "aiven_organizational_unit" "example_unit" {
  name      = "Example organizational unit - changed name at 12:34"
  parent_id = aiven_organization.main.org_id
}