// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package pcs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_pcs_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_pcs_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::PCS::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique Amazon Resource Name (ARN) of the cluster.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique Amazon Resource Name (ARN) of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of endpoints available for interaction with the scheduler.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An endpoint available for interaction with the scheduler.",
		//	    "properties": {
		//	      "Port": {
		//	        "description": "The endpoint's connection port number.",
		//	        "type": "string"
		//	      },
		//	      "PrivateIpAddress": {
		//	        "description": "The endpoint's private IP address.",
		//	        "type": "string"
		//	      },
		//	      "PublicIpAddress": {
		//	        "description": "The endpoint's public IP address.",
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "description": "Indicates the type of endpoint running at the specific IP address.",
		//	        "enum": [
		//	          "SLURMCTLD",
		//	          "SLURMDBD"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Port",
		//	      "PrivateIpAddress",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Port
					"port": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The endpoint's connection port number.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: PrivateIpAddress
					"private_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The endpoint's private IP address.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: PublicIpAddress
					"public_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The endpoint's public IP address.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Indicates the type of endpoint running at the specific IP address.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of endpoints available for interaction with the scheduler.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ErrorInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of errors that occurred during cluster provisioning.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An error that occurred during resource provisioning.",
		//	    "properties": {
		//	      "Code": {
		//	        "description": "The short-form error code.",
		//	        "type": "string"
		//	      },
		//	      "Message": {
		//	        "description": "The detailed error information.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"error_info": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Code
					"code": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The short-form error code.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Message
					"message": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The detailed error information.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of errors that occurred during cluster provisioning.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The generated unique ID of the cluster.",
		//	  "pattern": "^(pcs_[a-zA-Z0-9]+|[A-Za-z][A-Za-z0-9-]{1,40})$",
		//	  "type": "string"
		//	}
		"cluster_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The generated unique ID of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name that identifies the cluster.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name that identifies the cluster.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Networking
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The networking configuration for the cluster's control plane.",
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "description": "The list of security group IDs associated with the Elastic Network Interface (ENI) created in subnets.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "A VPC security group ID.",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SubnetIds": {
		//	      "description": "The list of subnet IDs where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources. The subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or an AWS Local Zone. AWS PCS currently supports only 1 subnet in this list.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "A VPC subnet ID.",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"networking": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of security group IDs associated with the Elastic Network Interface (ENI) created in subnets.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The list of subnet IDs where AWS PCS creates an Elastic Network Interface (ENI) to enable communication between managed controllers and AWS PCS resources. The subnet must have an available IP address, cannot reside in AWS Outposts, AWS Wavelength, or an AWS Local Zone. AWS PCS currently supports only 1 subnet in this list.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The networking configuration for the cluster's control plane.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scheduler
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The cluster management and job scheduling software associated with the cluster.",
		//	  "properties": {
		//	    "Type": {
		//	      "description": "The software AWS PCS uses to manage cluster scaling and job scheduling.",
		//	      "enum": [
		//	        "SLURM"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "description": "The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type",
		//	    "Version"
		//	  ],
		//	  "type": "object"
		//	}
		"scheduler": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The software AWS PCS uses to manage cluster scaling and job scheduling.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"SLURM",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The cluster management and job scheduling software associated with the cluster.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Size
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The size of the cluster.",
		//	  "enum": [
		//	    "SMALL",
		//	    "MEDIUM",
		//	    "LARGE"
		//	  ],
		//	  "type": "string"
		//	}
		"size": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The size of the cluster.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SMALL",
					"MEDIUM",
					"LARGE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SlurmConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Additional options related to the Slurm scheduler.",
		//	  "properties": {
		//	    "AuthKey": {
		//	      "additionalProperties": false,
		//	      "description": "The shared Slurm key for authentication, also known as the cluster secret.",
		//	      "properties": {
		//	        "SecretArn": {
		//	          "description": "The Amazon Resource Name (ARN) of the the shared Slurm key.",
		//	          "type": "string"
		//	        },
		//	        "SecretVersion": {
		//	          "description": "The version of the shared Slurm key.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "SecretArn",
		//	        "SecretVersion"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ScaleDownIdleTimeInSeconds": {
		//	      "description": "The time before an idle node is scaled down.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "SlurmCustomSettings": {
		//	      "description": "Additional Slurm-specific configuration that directly maps to Slurm settings.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Additional settings that directly map to Slurm settings.",
		//	        "properties": {
		//	          "ParameterName": {
		//	            "description": "AWS PCS supports configuration of the following Slurm parameters for clusters: Prolog, Epilog, and SelectTypeParameters.",
		//	            "type": "string"
		//	          },
		//	          "ParameterValue": {
		//	            "description": "The value for the configured Slurm setting.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ParameterName",
		//	          "ParameterValue"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"slurm_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthKey
				"auth_key": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SecretArn
						"secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the the shared Slurm key.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SecretVersion
						"secret_version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The version of the shared Slurm key.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The shared Slurm key for authentication, also known as the cluster secret.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ScaleDownIdleTimeInSeconds
				"scale_down_idle_time_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The time before an idle node is scaled down.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SlurmCustomSettings
				"slurm_custom_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ParameterName
							"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "AWS PCS supports configuration of the following Slurm parameters for clusters: Prolog, Epilog, and SelectTypeParameters.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: ParameterValue
							"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the configured Slurm setting.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Additional Slurm-specific configuration that directly maps to Slurm settings.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Additional options related to the Slurm scheduler.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The provisioning status of the cluster. The provisioning status doesn't indicate the overall health of the cluster.",
		//	  "enum": [
		//	    "CREATING",
		//	    "ACTIVE",
		//	    "UPDATING",
		//	    "DELETING",
		//	    "CREATE_FAILED",
		//	    "DELETE_FAILED",
		//	    "UPDATE_FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The provisioning status of the cluster. The provisioning status doesn't indicate the overall health of the cluster.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  }
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "1 or more tags added to the resource. Each tag consists of a tag key and tag value. The tag value is optional and can be an empty string.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "AWS::PCS::Cluster resource creates an AWS PCS cluster.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCS::Cluster").WithTerraformTypeName("awscc_pcs_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                             "Arn",
		"auth_key":                        "AuthKey",
		"cluster_id":                      "Id",
		"code":                            "Code",
		"endpoints":                       "Endpoints",
		"error_info":                      "ErrorInfo",
		"message":                         "Message",
		"name":                            "Name",
		"networking":                      "Networking",
		"parameter_name":                  "ParameterName",
		"parameter_value":                 "ParameterValue",
		"port":                            "Port",
		"private_ip_address":              "PrivateIpAddress",
		"public_ip_address":               "PublicIpAddress",
		"scale_down_idle_time_in_seconds": "ScaleDownIdleTimeInSeconds",
		"scheduler":                       "Scheduler",
		"secret_arn":                      "SecretArn",
		"secret_version":                  "SecretVersion",
		"security_group_ids":              "SecurityGroupIds",
		"size":                            "Size",
		"slurm_configuration":             "SlurmConfiguration",
		"slurm_custom_settings":           "SlurmCustomSettings",
		"status":                          "Status",
		"subnet_ids":                      "SubnetIds",
		"tags":                            "Tags",
		"type":                            "Type",
		"version":                         "Version",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(60)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}