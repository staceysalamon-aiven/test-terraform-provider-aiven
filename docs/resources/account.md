---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_account Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  Creates and manages an Aiven account.
  ~> This resource is deprecated
  This resource will be removed in v5.0.0. Use aiven_organization instead.
---

# aiven_account (Resource)

Creates and manages an Aiven account.

~> **This resource is deprecated**
This resource will be removed in v5.0.0. Use `aiven_organization` instead.

## Example Usage

```terraform
resource "aiven_account" "account1" {
  name = "<ACCOUNT_NAME>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Account name

### Optional

- `primary_billing_group_id` (String, Deprecated) Billing group id
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_id` (String, Deprecated) Account id
- `create_time` (String) Time of creation
- `id` (String) The ID of this resource.
- `is_account_owner` (Boolean, Deprecated) If true, user is part of the owners team for this account
- `owner_team_id` (String, Deprecated) Owner team id
- `tenant_id` (String) Tenant id
- `update_time` (String) Time of last update

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `default` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_account.account1 account_id
```
