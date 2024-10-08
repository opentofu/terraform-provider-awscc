// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	registry.AddResourceFactory("awscc_nimblestudio_studio", studioResource)
}

// studioResource returns the Terraform awscc_nimblestudio_studio resource.
// This Terraform resource corresponds to the CloudFormation AWS::NimbleStudio::Studio resource.
func studioResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AdminRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"admin_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.</p>",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA friendly name for the studio.\u003c/p\u003e",
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>A friendly name for the studio.</p>",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 64),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: HomeRegion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Web Services Region where the studio resource is located.\u003c/p\u003e",
		//	  "maxLength": 50,
		//	  "minLength": 0,
		//	  "pattern": "[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]",
		//	  "type": "string"
		//	}
		"home_region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Web Services Region where the studio resource is located.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SsoClientId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"sso_client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StudioEncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eConfiguration of the encryption method that is used for the studio.\u003c/p\u003e",
		//	  "properties": {
		//	    "KeyArn": {
		//	      "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
		//	      "minLength": 4,
		//	      "pattern": "^arn:.*",
		//	      "type": "string"
		//	    },
		//	    "KeyType": {
		//	      "description": "\u003cp\u003eThe type of KMS key that is used to encrypt studio data.\u003c/p\u003e",
		//	      "enum": [
		//	        "AWS_OWNED_KEY",
		//	        "CUSTOMER_MANAGED_KEY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "KeyType"
		//	  ],
		//	  "type": "object"
		//	}
		"studio_encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KeyArn
				"key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The ARN for a KMS key that is used to encrypt studio data.</p>",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtLeast(4),
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:.*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KeyType
				"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The type of KMS key that is used to encrypt studio data.</p>",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"AWS_OWNED_KEY",
							"CUSTOMER_MANAGED_KEY",
						),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>Configuration of the encryption method that is used for the studio.</p>",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StudioId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"studio_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StudioName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.\u003c/p\u003e",
		//	  "maxLength": 64,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9]*$",
		//	  "type": "string"
		//	}
		"studio_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.</p>",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StudioUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe address of the web page for the studio.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"studio_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The address of the web page for the studio.</p>",
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
		//	  "description": "",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe IAM role that Studio Users will assume when logging in to the Nimble Studio portal.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"user_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.</p>",
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
		Description: "Represents a studio that contains other Nimble Studio resources",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::Studio").WithTerraformTypeName("awscc_nimblestudio_studio")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"admin_role_arn":                  "AdminRoleArn",
		"display_name":                    "DisplayName",
		"home_region":                     "HomeRegion",
		"key_arn":                         "KeyArn",
		"key_type":                        "KeyType",
		"sso_client_id":                   "SsoClientId",
		"studio_encryption_configuration": "StudioEncryptionConfiguration",
		"studio_id":                       "StudioId",
		"studio_name":                     "StudioName",
		"studio_url":                      "StudioUrl",
		"tags":                            "Tags",
		"user_role_arn":                   "UserRoleArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
