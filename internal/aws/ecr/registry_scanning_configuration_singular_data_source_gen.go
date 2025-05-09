// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecr_registry_scanning_configuration", registryScanningConfigurationDataSource)
}

// registryScanningConfigurationDataSource returns the Terraform awscc_ecr_registry_scanning_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECR::RegistryScanningConfiguration resource.
func registryScanningConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: RegistryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The registry id.",
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The registry id.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Rules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scanning rules associated with the registry. A registry scanning configuration may contain a maximum of 2 rules.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A rule representing the details of a scanning configuration.",
		//	    "properties": {
		//	      "RepositoryFilters": {
		//	        "description": "The repository filters associated with the scanning configuration for a private registry.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The details of a scanning repository filter.",
		//	          "properties": {
		//	            "Filter": {
		//	              "description": "The filter to use when scanning.",
		//	              "pattern": "^[a-z0-9*](?:[._\\-/a-z0-9*]?[a-z0-9*]+)*$",
		//	              "type": "string"
		//	            },
		//	            "FilterType": {
		//	              "description": "The type associated with the filter.",
		//	              "enum": [
		//	                "WILDCARD"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Filter",
		//	            "FilterType"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "maxItems": 100,
		//	        "minItems": 0,
		//	        "type": "array"
		//	      },
		//	      "ScanFrequency": {
		//	        "description": "The frequency that scans are performed.",
		//	        "enum": [
		//	          "SCAN_ON_PUSH",
		//	          "CONTINUOUS_SCAN"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ScanFrequency",
		//	      "RepositoryFilters"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: RepositoryFilters
					"repository_filters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Filter
								"filter": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The filter to use when scanning.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: FilterType
								"filter_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The type associated with the filter.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The repository filters associated with the scanning configuration for a private registry.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ScanFrequency
					"scan_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The frequency that scans are performed.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The scanning rules associated with the registry. A registry scanning configuration may contain a maximum of 2 rules.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScanType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of scanning configured for the registry.",
		//	  "enum": [
		//	    "BASIC",
		//	    "ENHANCED"
		//	  ],
		//	  "type": "string"
		//	}
		"scan_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of scanning configured for the registry.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ECR::RegistryScanningConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::RegistryScanningConfiguration").WithTerraformTypeName("awscc_ecr_registry_scanning_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"filter":             "Filter",
		"filter_type":        "FilterType",
		"registry_id":        "RegistryId",
		"repository_filters": "RepositoryFilters",
		"rules":              "Rules",
		"scan_frequency":     "ScanFrequency",
		"scan_type":          "ScanType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
