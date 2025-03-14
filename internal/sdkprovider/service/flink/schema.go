// Code generated by user config generator. DO NOT EDIT.

package flink

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func flinkJarApplicationRename() map[string]string {
	return map[string]string{"application_id": "id"}
}

func flinkJarApplicationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Computed:    true,
			Description: "Application ID.",
			Type:        schema.TypeString,
		},
		"application_versions": {
			Computed:    true,
			Description: "JarApplicationVersions.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"created_at": {
					Computed:    true,
					Description: "The creation timestamp of this entity in ISO 8601 format, always in UTC.",
					Type:        schema.TypeString,
				},
				"created_by": {
					Computed:    true,
					Description: "The creator of this entity.",
					Type:        schema.TypeString,
				},
				"file_info": {
					Computed:    true,
					Description: "Flink JarApplicationVersion FileInfo.",
					Elem: &schema.Resource{Schema: map[string]*schema.Schema{
						"file_sha256": {
							Computed:    true,
							Description: "sha256 of the file if known.",
							Type:        schema.TypeString,
						},
						"file_size": {
							Computed:    true,
							Description: "The size of the file in bytes.",
							Type:        schema.TypeInt,
						},
						"file_status": {
							Computed:    true,
							Description: "Indicates whether the uploaded .jar file has been verified by the system and deployment ready. The possible values are `INITIAL`, `READY` and `FAILED`.",
							Type:        schema.TypeString,
						},
						"url": {
							Computed:    true,
							Description: "The pre-signed url of the bucket where the .jar file is uploaded. Becomes null when the JarApplicationVersion is ready or failed.",
							Type:        schema.TypeString,
						},
						"verify_error_code": {
							Computed:    true,
							Description: "In the case file_status is FAILED, the error code of the failure. The possible values are `1`, `2` and `3`.",
							Type:        schema.TypeInt,
						},
						"verify_error_message": {
							Computed:    true,
							Description: "In the case file_status is FAILED, may contain details about the failure.",
							Type:        schema.TypeString,
						},
					}},
					Type: schema.TypeList,
				},
				"id": {
					Computed:    true,
					Description: "ApplicationVersion ID.",
					Type:        schema.TypeString,
				},
				"version": {
					Computed:    true,
					Description: "Version number.",
					Type:        schema.TypeInt,
				},
			}},
			Type: schema.TypeList,
		},
		"created_at": {
			Computed:    true,
			Description: "The creation timestamp of this entity in ISO 8601 format, always in UTC.",
			Type:        schema.TypeString,
		},
		"created_by": {
			Computed:    true,
			Description: "The creator of this entity.",
			Type:        schema.TypeString,
		},
		"current_deployment": {
			Computed:    true,
			Description: "Flink JarApplicationDeployment.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"created_at": {
					Computed:    true,
					Description: "The creation timestamp of this entity in ISO 8601 format, always in UTC.",
					Type:        schema.TypeString,
				},
				"created_by": {
					Computed:    true,
					Description: "The creator of this entity.",
					Type:        schema.TypeString,
				},
				"entry_class": {
					Computed:    true,
					Description: "The fully qualified name of the entry class to pass during Flink job submission through the entryClass parameter.",
					Type:        schema.TypeString,
				},
				"error_msg": {
					Computed:    true,
					Description: "Error message describing what caused deployment to fail.",
					Type:        schema.TypeString,
				},
				"id": {
					Computed:    true,
					Description: "Deployment ID.",
					Type:        schema.TypeString,
				},
				"job_id": {
					Computed:    true,
					Description: "Job ID.",
					Type:        schema.TypeString,
				},
				"last_savepoint": {
					Computed:    true,
					Description: "Job savepoint.",
					Type:        schema.TypeString,
				},
				"parallelism": {
					Computed:    true,
					Description: "Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.",
					Type:        schema.TypeInt,
				},
				"program_args": {
					Computed:    true,
					Description: "Arguments to pass during Flink job submission through the programArgsList parameter.",
					Elem:        &schema.Schema{Type: schema.TypeString},
					Type:        schema.TypeSet,
				},
				"starting_savepoint": {
					Computed:    true,
					Description: "Job savepoint.",
					Type:        schema.TypeString,
				},
				"status": {
					Computed:    true,
					Description: "Deployment status. The possible values are `CANCELED`, `CANCELLING`, `CANCELLING_REQUESTED`, `CREATED`, `DELETE_REQUESTED`, `DELETING`, `FAILED`, `FAILING`, `FINISHED`, `INITIALIZING`, `RECONCILING`, `RESTARTING`, `RUNNING`, `SAVING`, `SAVING_AND_STOP`, `SAVING_AND_STOP_REQUESTED` and `SUSPENDED`.",
					Type:        schema.TypeString,
				},
				"version_id": {
					Computed:    true,
					Description: "ApplicationVersion ID.",
					Type:        schema.TypeString,
				},
			}},
			Type: schema.TypeList,
		},
		"name": {
			Description:  "Application name. Maximum length: `128`.",
			Required:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringLenBetween(0, 128),
		},
		"project": {
			Description: "Project name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"service_name": {
			Description: "Service name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"updated_at": {
			Computed:    true,
			Description: "The update timestamp of this entity in ISO 8601 format, always in UTC.",
			Type:        schema.TypeString,
		},
		"updated_by": {
			Computed:    true,
			Description: "The latest updater of this entity.",
			Type:        schema.TypeString,
		},
	}
}

func flinkJarApplicationDeploymentRename() map[string]string {
	return map[string]string{"deployment_id": "id"}
}

func flinkJarApplicationDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Description: "Application Id. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"created_at": {
			Computed:    true,
			Description: "The creation timestamp of this entity in ISO 8601 format, always in UTC.",
			Type:        schema.TypeString,
		},
		"created_by": {
			Computed:    true,
			Description: "The creator of this entity.",
			Type:        schema.TypeString,
		},
		"deployment_id": {
			Computed:    true,
			Description: "Deployment ID.",
			Type:        schema.TypeString,
		},
		"entry_class": {
			Computed:     true,
			Description:  "The fully qualified name of the entry class to pass during Flink job submission through the entryClass parameter. Maximum length: `128`.",
			Optional:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringLenBetween(1, 128),
		},
		"error_msg": {
			Computed:    true,
			Description: "Error message describing what caused deployment to fail.",
			Type:        schema.TypeString,
		},
		"job_id": {
			Computed:    true,
			Description: "Job ID.",
			Type:        schema.TypeString,
		},
		"last_savepoint": {
			Computed:    true,
			Description: "Job savepoint.",
			Type:        schema.TypeString,
		},
		"parallelism": {
			Computed:     true,
			Description:  "Reading of Flink parallel execution documentation is recommended before setting this value to other than 1. Please do not set this value higher than (total number of nodes x number_of_task_slots), or every new job created will fail.",
			Optional:     true,
			Type:         schema.TypeInt,
			ValidateFunc: validation.IntBetween(1, 128),
		},
		"program_args": {
			Computed:    true,
			Description: "Arguments to pass during Flink job submission through the programArgsList parameter.",
			Elem:        &schema.Schema{Type: schema.TypeString},
			MaxItems:    32,
			Optional:    true,
			Type:        schema.TypeSet,
		},
		"project": {
			Description: "Project name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"restart_enabled": {
			Description: "Specifies whether a Flink Job is restarted in case it fails. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Optional:    true,
			Type:        schema.TypeBool,
		},
		"service_name": {
			Description: "Service name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"starting_savepoint": {
			Computed:     true,
			Description:  "Job savepoint. Maximum length: `2048`.",
			Optional:     true,
			Type:         schema.TypeString,
			ValidateFunc: validation.StringLenBetween(1, 2048),
		},
		"status": {
			Computed:    true,
			Description: "Deployment status. The possible values are `CANCELED`, `CANCELLING`, `CANCELLING_REQUESTED`, `CREATED`, `DELETE_REQUESTED`, `DELETING`, `FAILED`, `FAILING`, `FINISHED`, `INITIALIZING`, `RECONCILING`, `RESTARTING`, `RUNNING`, `SAVING`, `SAVING_AND_STOP`, `SAVING_AND_STOP_REQUESTED` and `SUSPENDED`.",
			Type:        schema.TypeString,
		},
		"version_id": {
			Description: "ApplicationVersion ID. Maximum length: `36`. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
	}
}

func flinkJarApplicationVersionRename() map[string]string {
	return map[string]string{"application_version_id": "id"}
}

func flinkJarApplicationVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Description: "Application Id. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"application_version_id": {
			Computed:    true,
			Description: "ApplicationVersion ID.",
			Type:        schema.TypeString,
		},
		"created_at": {
			Computed:    true,
			Description: "The creation timestamp of this entity in ISO 8601 format, always in UTC.",
			Type:        schema.TypeString,
		},
		"created_by": {
			Computed:    true,
			Description: "The creator of this entity.",
			Type:        schema.TypeString,
		},
		"file_info": {
			Computed:    true,
			Description: "Flink JarApplicationVersion FileInfo.",
			Elem: &schema.Resource{Schema: map[string]*schema.Schema{
				"file_sha256": {
					Computed:    true,
					Description: "sha256 of the file if known.",
					Type:        schema.TypeString,
				},
				"file_size": {
					Computed:    true,
					Description: "The size of the file in bytes.",
					Type:        schema.TypeInt,
				},
				"file_status": {
					Computed:    true,
					Description: "Indicates whether the uploaded .jar file has been verified by the system and deployment ready. The possible values are `FAILED`, `INITIAL` and `READY`.",
					Type:        schema.TypeString,
				},
				"url": {
					Computed:    true,
					Description: "The pre-signed url of the bucket where the .jar file is uploaded. Becomes null when the JarApplicationVersion is ready or failed.",
					Type:        schema.TypeString,
				},
				"verify_error_code": {
					Computed:    true,
					Description: "In the case file_status is FAILED, the error code of the failure. The possible values are `1`, `2` and `3`.",
					Type:        schema.TypeInt,
				},
				"verify_error_message": {
					Computed:    true,
					Description: "In the case file_status is FAILED, may contain details about the failure.",
					Type:        schema.TypeString,
				},
			}},
			Type: schema.TypeList,
		},
		"project": {
			Description: "Project name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"service_name": {
			Description: "Service name. Changing this property forces recreation of the resource.",
			ForceNew:    true,
			Required:    true,
			Type:        schema.TypeString,
		},
		"version": {
			Computed:    true,
			Description: "Version number.",
			Type:        schema.TypeInt,
		},
	}
}
