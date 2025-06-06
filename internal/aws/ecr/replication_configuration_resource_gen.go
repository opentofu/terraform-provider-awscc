// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ecr_replication_configuration", replicationConfigurationResource)
}

// replicationConfigurationResource returns the Terraform awscc_ecr_replication_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECR::ReplicationConfiguration resource.
func replicationConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RegistryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The replication configuration for a registry.",
		//	  "properties": {
		//	    "Rules": {
		//	      "description": "An array of objects representing the replication destinations and repository filters for a replication configuration.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An array of objects representing the replication destinations and repository filters for a replication configuration.",
		//	        "properties": {
		//	          "Destinations": {
		//	            "description": "An array of objects representing the destination for a replication rule.",
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "An array of objects representing the destination for a replication rule.",
		//	              "properties": {
		//	                "Region": {
		//	                  "description": "The Region to replicate to.",
		//	                  "pattern": "[0-9a-z-]{2,25}",
		//	                  "type": "string"
		//	                },
		//	                "RegistryId": {
		//	                  "description": "The AWS account ID of the Amazon ECR private registry to replicate to. When configuring cross-Region replication within your own registry, specify your own account ID.",
		//	                  "pattern": "^[0-9]{12}$",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Region",
		//	                "RegistryId"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 1,
		//	            "type": "array"
		//	          },
		//	          "RepositoryFilters": {
		//	            "description": "An array of objects representing the filters for a replication rule. Specifying a repository filter for a replication rule provides a method for controlling which repositories in a private registry are replicated.",
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "The filter settings used with image replication. Specifying a repository filter to a replication rule provides a method for controlling which repositories in a private registry are replicated. If no filters are added, the contents of all repositories are replicated.",
		//	              "properties": {
		//	                "Filter": {
		//	                  "description": "The repository filter details. When the ``PREFIX_MATCH`` filter type is specified, this value is required and should be the repository name prefix to configure replication for.",
		//	                  "pattern": "^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$",
		//	                  "type": "string"
		//	                },
		//	                "FilterType": {
		//	                  "description": "The repository filter type. The only supported value is ``PREFIX_MATCH``, which is a repository name prefix specified with the ``filter`` parameter.",
		//	                  "enum": [
		//	                    "PREFIX_MATCH"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Filter",
		//	                "FilterType"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 100,
		//	            "minItems": 0,
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Destinations"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 10,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "Rules"
		//	  ],
		//	  "type": "object"
		//	}
		"replication_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Rules
				"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destinations
							"destinations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Region
										"region": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The Region to replicate to.",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-z-]{2,25}"), ""),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
										// Property: RegistryId
										"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The AWS account ID of the Amazon ECR private registry to replicate to. When configuring cross-Region replication within your own registry, specify your own account ID.",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of objects representing the destination for a replication rule.",
								Required:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeBetween(1, 100),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: RepositoryFilters
							"repository_filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Filter
										"filter": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The repository filter details. When the ``PREFIX_MATCH`` filter type is specified, this value is required and should be the repository name prefix to configure replication for.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.RegexMatches(regexp.MustCompile("^(?:[a-z0-9]+(?:[._-][a-z0-9]*)*/)*[a-z0-9]*(?:[._-][a-z0-9]*)*$"), ""),
												fwvalidators.NotNullString(),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: FilterType
										"filter_type": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The repository filter type. The only supported value is ``PREFIX_MATCH``, which is a repository name prefix specified with the ``filter`` parameter.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.OneOf(
													"PREFIX_MATCH",
												),
												fwvalidators.NotNullString(),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of objects representing the filters for a replication rule. Specifying a repository filter for a replication rule provides a method for controlling which repositories in a private registry are replicated.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.List{ /*START VALIDATORS*/
									listvalidator.SizeBetween(0, 100),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									listplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of objects representing the replication destinations and repository filters for a replication configuration.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 10),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The replication configuration for a registry.",
			Required:    true,
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
		Description: "The ``AWS::ECR::ReplicationConfiguration`` resource creates or updates the replication configuration for a private registry. The first time a replication configuration is applied to a private registry, a service-linked IAM role is created in your account for the replication process. For more information, see [Using Service-Linked Roles for Amazon ECR](https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html) in the *Amazon Elastic Container Registry User Guide*.\n  When configuring cross-account replication, the destination account must grant the source account permission to replicate. This permission is controlled using a private registry permissions policy. For more information, see ``AWS::ECR::RegistryPolicy``.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::ReplicationConfiguration").WithTerraformTypeName("awscc_ecr_replication_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destinations":              "Destinations",
		"filter":                    "Filter",
		"filter_type":               "FilterType",
		"region":                    "Region",
		"registry_id":               "RegistryId",
		"replication_configuration": "ReplicationConfiguration",
		"repository_filters":        "RepositoryFilters",
		"rules":                     "Rules",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
