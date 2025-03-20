resource "aiven_organizational_unit" "example_unit" {
  name      = "Example organizational unit change"
  parent_id = aiven_organization.main.org_id
}