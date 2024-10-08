// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package efs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_efs_access_point", accessPointResource)
}

// accessPointResource returns the Terraform awscc_efs_access_point resource.
// This Terraform resource corresponds to the CloudFormation AWS::EFS::AccessPoint resource.
func accessPointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessPointId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"access_point_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AccessPointTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag is a key-value pair attached to a file system. Allowed characters in the ``Key`` and ``Value`` properties are letters, white space, and numbers that can be represented in UTF-8, and the following characters:``+ - = . _ : /``",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key (String). The key can't start with ``aws:``.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag key.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"access_point_tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key (String). The key can't start with ``aws:``.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag key.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClientToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The opaque string specified in the request to ensure idempotent creation.",
		//	  "type": "string"
		//	}
		"client_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The opaque string specified in the request to ensure idempotent creation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the EFS file system that the access point applies to. Accepts only the ID format for input when specifying a file system, for example ``fs-0123456789abcedf2``.",
		//	  "type": "string"
		//	}
		"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the EFS file system that the access point applies to. Accepts only the ID format for input when specifying a file system, for example ``fs-0123456789abcedf2``.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PosixUser
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.",
		//	  "properties": {
		//	    "Gid": {
		//	      "description": "The POSIX group ID used for all file system operations using this access point.",
		//	      "type": "string"
		//	    },
		//	    "SecondaryGids": {
		//	      "description": "Secondary POSIX group IDs used for all file system operations using this access point.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Uid": {
		//	      "description": "The POSIX user ID used for all file system operations using this access point.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Uid",
		//	    "Gid"
		//	  ],
		//	  "type": "object"
		//	}
		"posix_user": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Gid
				"gid": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The POSIX group ID used for all file system operations using this access point.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SecondaryGids
				"secondary_gids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Secondary POSIX group IDs used for all file system operations using this access point.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
						listplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Uid
				"uid": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The POSIX user ID used for all file system operations using this access point.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The full POSIX identity, including the user ID, group ID, and secondary group IDs on the access point that is used for all file operations by NFS clients using the access point.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RootDirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The directory on the EFS file system that the access point exposes as the root directory to NFS clients using the access point.",
		//	  "properties": {
		//	    "CreationInfo": {
		//	      "additionalProperties": false,
		//	      "description": "(Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` \u003e ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. \n  If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` \u003e ``Path`` does not exist, attempts to mount the file system using the access point will fail.",
		//	      "properties": {
		//	        "OwnerGid": {
		//	          "description": "Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).",
		//	          "type": "string"
		//	        },
		//	        "OwnerUid": {
		//	          "description": "Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).",
		//	          "type": "string"
		//	        },
		//	        "Permissions": {
		//	          "description": "Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.",
		//	          "pattern": "^[0-7]{3,4}$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "OwnerUid",
		//	        "OwnerGid",
		//	        "Permissions"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Path": {
		//	      "description": "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.",
		//	      "maxLength": 100,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"root_directory": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CreationInfo
				"creation_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OwnerGid
						"owner_gid": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies the POSIX group ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: OwnerUid
						"owner_uid": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies the POSIX user ID to apply to the ``RootDirectory``. Accepts values from 0 to 2^32 (4294967295).",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Permissions
						"permissions": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Specifies the POSIX permissions to apply to the ``RootDirectory``, in the format of an octal number representing the file's mode bits.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^[0-7]{3,4}$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "(Optional) Specifies the POSIX IDs and permissions to apply to the access point's ``RootDirectory``. If the ``RootDirectory`` > ``Path`` specified does not exist, EFS creates the root directory using the ``CreationInfo`` settings when a client connects to an access point. When specifying the ``CreationInfo``, you must provide values for all properties. \n  If you do not provide ``CreationInfo`` and the specified ``RootDirectory`` > ``Path`` does not exist, attempts to mount the file system using the access point will fail.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
						objectplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Path
				"path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide the ``CreationInfo``.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The directory on the EFS file system that the access point exposes as the root directory to NFS clients using the access point.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "The ``AWS::EFS::AccessPoint`` resource creates an EFS access point. An access point is an application-specific view into an EFS file system that applies an operating system user and group, and a file system path, to any file system request made through the access point. The operating system user and group override any identity information provided by the NFS client. The file system path is exposed as the access point's root directory. Applications using the access point can only access data in its own directory and below. To learn more, see [Mounting a file system using EFS access points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).\n This operation requires permissions for the ``elasticfilesystem:CreateAccessPoint`` action.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EFS::AccessPoint").WithTerraformTypeName("awscc_efs_access_point")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_point_id":   "AccessPointId",
		"access_point_tags": "AccessPointTags",
		"arn":               "Arn",
		"client_token":      "ClientToken",
		"creation_info":     "CreationInfo",
		"file_system_id":    "FileSystemId",
		"gid":               "Gid",
		"key":               "Key",
		"owner_gid":         "OwnerGid",
		"owner_uid":         "OwnerUid",
		"path":              "Path",
		"permissions":       "Permissions",
		"posix_user":        "PosixUser",
		"root_directory":    "RootDirectory",
		"secondary_gids":    "SecondaryGids",
		"uid":               "Uid",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
