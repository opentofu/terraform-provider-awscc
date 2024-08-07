// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_secretsmanager_resource_policy", resourcePolicyDataSource)
}

// resourcePolicyDataSource returns the Terraform awscc_secretsmanager_resource_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecretsManager::ResourcePolicy resource.
func resourcePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BlockPublicPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to block resource-based policies that allow broad access to the secret.",
		//	  "type": "boolean"
		//	}
		"block_public_policy": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to block resource-based policies that allow broad access to the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Arn of the secret.",
		//	  "type": "string"
		//	}
		"resource_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Arn of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourcePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A JSON-formatted string for an AWS resource-based policy.",
		//	  "type": "string"
		//	}
		"resource_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A JSON-formatted string for an AWS resource-based policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN or name of the secret to attach the resource-based policy.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN or name of the secret to attach the resource-based policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecretsManager::ResourcePolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecretsManager::ResourcePolicy").WithTerraformTypeName("awscc_secretsmanager_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"block_public_policy": "BlockPublicPolicy",
		"resource_policy":     "ResourcePolicy",
		"resource_policy_id":  "Id",
		"secret_id":           "SecretId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
