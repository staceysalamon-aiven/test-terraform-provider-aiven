package organization

import (
	"context"

	avngen "github.com/aiven/go-client-codegen"
	"github.com/aiven/go-client-codegen/handler/account"
	"github.com/aiven/go-client-codegen/handler/organization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

var aivenOrganizationalPermissionSchema = map[string]*schema.Schema{
	"organization_id": {
		Type:        schema.TypeString,
		Description: "Organization ID",
		Required:    true,
	},
	"resource_type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: userconfig.Desc("Resource type.").PossibleValuesString("project").Build(),
	},
	"resource_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Resource Id.",
	},
	"permissions": {
		Type:        schema.TypeSet,
		Description: "A permission to set",
		Required:    true,
		Elem: &schema.Resource{
			Schema: permissionFields,
		},
	},
}

var permissionFields = map[string]*schema.Schema{
	"principal_type": {
		Type:        schema.TypeString,
		Required:    true,
		Description: userconfig.Desc("Type of the principal.").PossibleValuesString(organization.PrincipalTypeChoices()...).Build(),
	},
	"principal_id": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "ID of the principal.",
	},
	"permissions": {
		Type:        schema.TypeSet,
		Description: userconfig.Desc("List of permissions").PossibleValuesString(account.MemberTypeChoices()...).Build(),
		Required:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
	},
	"create_time": {
		Type:        schema.TypeString,
		Description: "Create Time",
		Computed:    true,
	},
	"update_time": {
		Type:        schema.TypeString,
		Description: "Update Time",
		Computed:    true,
	},
}

func ResourceOrganizationalPermission() *schema.Resource {
	return &schema.Resource{
		Description:   "Grants permissions to a principal for a resource.",
		CreateContext: common.WithGenClient(resourceOrganizationalPermissionUpsert),
		ReadContext:   common.WithGenClient(resourceOrganizationalPermissionRead),
		UpdateContext: common.WithGenClient(resourceOrganizationalPermissionUpsert),
		DeleteContext: common.WithGenClient(resourceOrganizationalPermissionDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),
		Schema:   aivenOrganizationalPermissionSchema,
	}
}

func resourceOrganizationalPermissionUpsert(ctx context.Context, d *schema.ResourceData, client avngen.Client) error {
	orgID := d.Get("organization_id").(string)
	resourceType := d.Get("resource_type").(string)
	resourceID := d.Get("resource_id").(string)

	req := new(organization.PermissionsUpdateIn)
	err := schemautil.ResourceDataGet(d, req)
	if err != nil {
		return err
	}

	// Must remove all existing permissions by sending an empty list per type and id
	old, err := client.PermissionsGet(ctx, orgID, organization.ResourceType(resourceType), resourceID)
	if err != nil && !avngen.IsNotFound(err) {
		return err
	}

outer:
	for _, o := range old {
		for _, n := range req.Permissions {
			if n.PrincipalType == o.PrincipalType && o.PrincipalId == n.PrincipalId {
				continue outer
			}
		}

		n := organization.PermissionIn{
			PrincipalType: o.PrincipalType,
			PrincipalId:   o.PrincipalId,
			Permissions:   make([]string, 0),
		}
		req.Permissions = append(req.Permissions, n)
	}

	err = client.PermissionsUpdate(ctx, orgID, organization.ResourceType(resourceType), resourceID, req)
	if err != nil {
		return err
	}

	d.SetId(schemautil.BuildResourceID(orgID, resourceType, resourceID))
	return resourceOrganizationalPermissionRead(ctx, d, client)
}

func resourceOrganizationalPermissionRead(ctx context.Context, d *schema.ResourceData, client avngen.Client) error {
	orgID, resourceType, resourceID, err := schemautil.SplitResourceID3(d.Id())
	if err != nil {
		return err
	}

	out, err := client.PermissionsGet(ctx, orgID, organization.ResourceType(resourceType), resourceID)
	if err != nil {
		return err
	}

	permissions := make([]map[string]any, 0, len(out))
	err = schemautil.Remarshal(out, &permissions)
	if err != nil {
		return err
	}

	// Removes fields that are not on the schema,
	// so it won't blow up when the DTO gets new fields with the updates
	for _, m := range permissions {
		for k := range m {
			if _, ok := permissionFields[k]; !ok {
				delete(m, k)
			}
		}
	}

	return d.Set("permissions", permissions)
}

func resourceOrganizationalPermissionDelete(ctx context.Context, d *schema.ResourceData, client avngen.Client) error {
	err := d.Set("permissions", []any{})
	if err != nil {
		return err
	}
	return resourceOrganizationalPermissionUpsert(ctx, d, client)
}
