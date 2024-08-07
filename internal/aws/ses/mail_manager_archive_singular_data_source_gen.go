// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ses_mail_manager_archive", mailManagerArchiveDataSource)
}

// mailManagerArchiveDataSource returns the Terraform awscc_ses_mail_manager_archive data source.
// This Terraform data source corresponds to the CloudFormation AWS::SES::MailManagerArchive resource.
func mailManagerArchiveDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ArchiveArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"archive_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ArchiveId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 66,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"archive_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ArchiveName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9_-]*[a-zA-Z0-9]$",
		//	  "type": "string"
		//	}
		"archive_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ArchiveState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ACTIVE",
		//	    "PENDING_DELETION"
		//	  ],
		//	  "type": "string"
		//	}
		"archive_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:aws(|-cn|-us-gov):kms:[a-z0-9-]{1,20}:[0-9]{12}:(key|alias)/.+$",
		//	  "type": "string"
		//	}
		"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Retention
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "RetentionPeriod": {
		//	      "enum": [
		//	        "THREE_MONTHS",
		//	        "SIX_MONTHS",
		//	        "NINE_MONTHS",
		//	        "ONE_YEAR",
		//	        "EIGHTEEN_MONTHS",
		//	        "TWO_YEARS",
		//	        "THIRTY_MONTHS",
		//	        "THREE_YEARS",
		//	        "FOUR_YEARS",
		//	        "FIVE_YEARS",
		//	        "SIX_YEARS",
		//	        "SEVEN_YEARS",
		//	        "EIGHT_YEARS",
		//	        "NINE_YEARS",
		//	        "TEN_YEARS",
		//	        "PERMANENT"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"retention": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RetentionPeriod
				"retention_period": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SES::MailManagerArchive",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::MailManagerArchive").WithTerraformTypeName("awscc_ses_mail_manager_archive")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"archive_arn":      "ArchiveArn",
		"archive_id":       "ArchiveId",
		"archive_name":     "ArchiveName",
		"archive_state":    "ArchiveState",
		"key":              "Key",
		"kms_key_arn":      "KmsKeyArn",
		"retention":        "Retention",
		"retention_period": "RetentionPeriod",
		"tags":             "Tags",
		"value":            "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
