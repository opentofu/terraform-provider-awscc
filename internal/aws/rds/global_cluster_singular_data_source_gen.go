// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_global_cluster", globalClusterDataSource)
}

// globalClusterDataSource returns the Terraform awscc_rds_global_cluster data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::GlobalCluster resource.
func globalClusterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeletionProtection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
		//	  "type": "boolean"
		//	}
		"deletion_protection": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "enum": [
		//	    "aurora",
		//	    "aurora-mysql",
		//	    "aurora-postgresql"
		//	  ],
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineLifecycleSupport
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The life cycle type of the global cluster. You can use this setting to enroll your global cluster into Amazon RDS Extended Support.",
		//	  "type": "string"
		//	}
		"engine_lifecycle_support": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The life cycle type of the global cluster. You can use this setting to enroll your global cluster into Amazon RDS Extended Support.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z]{1}(?:-?[a-zA-Z0-9]){0,62}$",
		//	  "type": "string"
		//	}
		"global_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalEndpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Address": {
		//	      "description": "The writer endpoint for the global database cluster. This endpoint always points to the writer DB instance in the current primary cluster.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"global_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Address
				"address": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The writer endpoint for the global database cluster. This endpoint always points to the writer DB instance in the current primary cluster.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SourceDBClusterIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
		//	  "oneOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "type": "string"
		//	}
		"source_db_cluster_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageEncrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
		//	  "type": "boolean"
		//	}
		"storage_encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RDS::GlobalCluster",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::GlobalCluster").WithTerraformTypeName("awscc_rds_global_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                      "Address",
		"deletion_protection":          "DeletionProtection",
		"engine":                       "Engine",
		"engine_lifecycle_support":     "EngineLifecycleSupport",
		"engine_version":               "EngineVersion",
		"global_cluster_identifier":    "GlobalClusterIdentifier",
		"global_endpoint":              "GlobalEndpoint",
		"key":                          "Key",
		"source_db_cluster_identifier": "SourceDBClusterIdentifier",
		"storage_encrypted":            "StorageEncrypted",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
