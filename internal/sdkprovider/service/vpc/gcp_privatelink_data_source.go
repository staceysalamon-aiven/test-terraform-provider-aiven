package vpc

import (
	"context"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceGCPPrivatelink() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceGCPPrivatelinkRead,
		Description: "The GCP Privatelink resource allows the creation and management of Aiven GCP Privatelink for a services.",
		Schema:      schemautil.ResourceSchemaAsDatasourceSchema(aivenGCPPrivatelinkSchema, "project", "service_name"),
	}
}

func datasourceGCPPrivatelinkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	projectName := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)
	d.SetId(schemautil.BuildResourceID(projectName, serviceName))

	return resourceGCPPrivatelinkRead(ctx, d, m)
}
