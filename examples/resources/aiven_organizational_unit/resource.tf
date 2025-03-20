resource "aiven_organizational_unit" "example_unit" {
  name      = "Example organizational unit - changed name at 16:08"
  parent_id = aiven_organization.main.org_id
}