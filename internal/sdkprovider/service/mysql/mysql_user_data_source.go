package mysql

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func DatasourceMySQLUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: schemautil.DatasourceServiceUserRead,
		Description: "The MySQL User data source provides information about the existing Aiven MySQL User.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(aivenMySQLUserSchema,
			"project", "service_name", "username"),
	}
}
