resource "aiven_organizational_unit" "example_unit" {
  name      = "0944 changed name"
  parent_id = aiven_organization.main.org_id
}