// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datazone_domain_unit", domainUnitDataSource)
}

// domainUnitDataSource returns the Terraform awscc_datazone_domain_unit data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataZone::DomainUnit resource.
func domainUnitDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which the domain unit was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp at which the domain unit was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the domain unit.",
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain where the domain unit was created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain where the domain unit was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain where you want to create a domain unit.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain where you want to create a domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"domain_unit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Identifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the domain unit that you want to get.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the domain unit that you want to get.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which the domain unit was last updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp at which the domain unit was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the domain unit.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w -]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ParentDomainUnitId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the parent domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"parent_domain_unit_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the parent domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ParentDomainUnitIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the parent domain unit.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-z0-9_-]+$",
		//	  "type": "string"
		//	}
		"parent_domain_unit_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the parent domain unit.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataZone::DomainUnit",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::DomainUnit").WithTerraformTypeName("awscc_datazone_domain_unit")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":                    "CreatedAt",
		"description":                   "Description",
		"domain_id":                     "DomainId",
		"domain_identifier":             "DomainIdentifier",
		"domain_unit_id":                "Id",
		"identifier":                    "Identifier",
		"last_updated_at":               "LastUpdatedAt",
		"name":                          "Name",
		"parent_domain_unit_id":         "ParentDomainUnitId",
		"parent_domain_unit_identifier": "ParentDomainUnitIdentifier",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
