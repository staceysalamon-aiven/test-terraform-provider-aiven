// Code generated by user config generator. DO NOT EDIT.

package serviceintegration

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/diff"
)

func metricsUserConfig() *schema.Schema {
	return &schema.Schema{
		Description:      "Metrics user configurable settings. **Warning:** There's no way to reset advanced configuration options to default. Options that you add cannot be removed later",
		DiffSuppressFunc: diff.SuppressUnchanged,
		Elem: &schema.Resource{Schema: map[string]*schema.Schema{
			"database": {
				Description: "Name of the database where to store metric datapoints. Only affects PostgreSQL destinations. Defaults to `metrics`. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"retention_days": {
				Description: "Number of days to keep old metrics. Only affects PostgreSQL destinations. Set to 0 for no automatic cleanup. Defaults to 30 days.",
				Optional:    true,
				Type:        schema.TypeInt,
			},
			"ro_username": {
				Description: "Name of a user that can be used to read metrics. This will be used for Grafana integration (if enabled) to prevent Grafana users from making undesired changes. Only affects PostgreSQL destinations. Defaults to `metrics_reader`. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"source_mysql": {
				Description: "Configuration options for metrics where source service is MySQL",
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"telegraf": {
					Description: "Configuration options for Telegraf MySQL input plugin",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"gather_event_waits": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_file_events_stats": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_index_io_waits": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_info_schema_auto_inc": {
							Description: "Gather auto_increment columns and max values from information schema.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_innodb_metrics": {
							Description: "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_perf_events_statements": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_process_list": {
							Description: "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_slave_status": {
							Description: "Gather metrics from SHOW SLAVE STATUS command output.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_table_io_waits": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_table_lock_waits": {
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"gather_table_schema": {
							Description: "Gather metrics from INFORMATION_SCHEMA.TABLES.",
							Optional:    true,
							Type:        schema.TypeBool,
						},
						"perf_events_statements_digest_text_limit": {
							Description: "Truncates digest text from perf_events_statements into this many characters. Example: `120`.",
							Optional:    true,
							Type:        schema.TypeInt,
						},
						"perf_events_statements_limit": {
							Description: "Limits metrics from perf_events_statements. Example: `250`.",
							Optional:    true,
							Type:        schema.TypeInt,
						},
						"perf_events_statements_time_limit": {
							Description: "Only include perf_events_statements whose last seen is less than this many seconds. Example: `86400`.",
							Optional:    true,
							Type:        schema.TypeInt,
						},
					}},
					MaxItems: 1,
					Optional: true,
					Type:     schema.TypeList,
				}}},
				MaxItems: 1,
				Optional: true,
				Type:     schema.TypeList,
			},
			"username": {
				Description: "Name of the user used to write metrics. Only affects PostgreSQL destinations. Defaults to `metrics_writer`. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				Optional:    true,
				Type:        schema.TypeString,
			},
		}},
		MaxItems: 1,
		Optional: true,
		Type:     schema.TypeList,
	}
}
