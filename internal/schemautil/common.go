package schemautil

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

//goland:noinspection GoDeprecation
func GetACLUserValidateFunc() schema.SchemaValidateFunc { //nolint:staticcheck
	return validation.StringMatch(
		regexp.MustCompile(`^[-._*?A-Za-z0-9]+$`),
		"Must consist of alpha-numeric characters, underscores, dashes, dots, and glob characters '*' and '?'")
}

//goland:noinspection GoDeprecation
func GetServiceUserValidateFunc() schema.SchemaValidateFunc { //nolint:staticcheck
	return validation.StringMatch(
		regexp.MustCompile(`^(\*$|[a-zA-Z0-9_?][a-zA-Z0-9-_?*.].{0,62})$`),
		"username should be alphanumeric, may not start with dash or dot, max 64 characters")
}

var (
	CommonSchemaProjectReference = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]*$"), "project name should be alphanumeric"),
		Description:  userconfig.Desc("The name of the project this resource belongs to.").ForceNew().Referenced().Build(),
	}

	CommonSchemaServiceNameReference = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_-]*$"), "common name should be alphanumeric"),
		Description:  userconfig.Desc("The name of the service that this resource belongs to.").ForceNew().Referenced().Build(),
	}
)

func SetTagsTerraformProperties(t map[string]string) []map[string]interface{} {
	tags := make([]map[string]interface{}, len(t))
	var i int
	for k, v := range t {
		tags[i] = map[string]interface{}{
			"key":   k,
			"value": v,
		}
		i++
	}

	return tags
}

func GetTagsFromSchema(d ResourceData) map[string]string {
	tags := make(map[string]string)

	for _, tag := range d.Get("tag").(*schema.Set).List() {
		tagVal := tag.(map[string]interface{})
		tags[tagVal["key"].(string)] = tagVal["value"].(string)
	}

	return tags
}

// PointerValueOrDefault returns pointer's value or default
func PointerValueOrDefault[T comparable](v *T, d T) T {
	if v == nil {
		return d
	}
	return *v
}

// ComposeContexts composes multiple context (create, update, read or delete) functions into one.
// So instead of chaining them, they can be composed
func ComposeContexts(funcs ...func(context.Context, *schema.ResourceData, any) diag.Diagnostics) func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
		for _, f := range funcs {
			err := f(ctx, d, m)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
