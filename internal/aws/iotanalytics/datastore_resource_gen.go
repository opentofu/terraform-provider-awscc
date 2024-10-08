// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotanalytics

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_iotanalytics_datastore", datastoreResource)
}

// datastoreResource returns the Terraform awscc_iotanalytics_datastore resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTAnalytics::Datastore resource.
func datastoreResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatastoreName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"datastore_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastorePartitions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Partitions": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Partition": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "AttributeName": {
		//	                "pattern": "[a-zA-Z0-9_]+",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "AttributeName"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "TimestampPartition": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "AttributeName": {
		//	                "pattern": "[a-zA-Z0-9_]+",
		//	                "type": "string"
		//	              },
		//	              "TimestampFormat": {
		//	                "pattern": "[a-zA-Z0-9\\s\\[\\]_,.'/:-]*",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "AttributeName"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 25,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"datastore_partitions": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Partitions
				"partitions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Partition
							"partition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AttributeName
									"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_]+"), ""),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: TimestampPartition
							"timestamp_partition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AttributeName
									"attribute_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_]+"), ""),
											fwvalidators.NotNullString(),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: TimestampFormat
									"timestamp_format": schema.StringAttribute{ /*START ATTRIBUTE*/
										Optional: true,
										Computed: true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9\\s\\[\\]_,.'/:-]*"), ""),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 25),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DatastoreStorage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CustomerManagedS3": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Bucket": {
		//	          "maxLength": 255,
		//	          "minLength": 3,
		//	          "pattern": "[a-zA-Z0-9.\\-_]*",
		//	          "type": "string"
		//	        },
		//	        "KeyPrefix": {
		//	          "maxLength": 255,
		//	          "minLength": 1,
		//	          "pattern": "[a-zA-Z0-9!_.*'()/{}:-]*/",
		//	          "type": "string"
		//	        },
		//	        "RoleArn": {
		//	          "maxLength": 2048,
		//	          "minLength": 20,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Bucket",
		//	        "RoleArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IotSiteWiseMultiLayerStorage": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CustomerManagedS3Storage": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Bucket": {
		//	              "maxLength": 255,
		//	              "minLength": 3,
		//	              "pattern": "[a-zA-Z0-9.\\-_]*",
		//	              "type": "string"
		//	            },
		//	            "KeyPrefix": {
		//	              "maxLength": 255,
		//	              "minLength": 1,
		//	              "pattern": "[a-zA-Z0-9!_.*'()/{}:-]*/",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Bucket"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ServiceManagedS3": {
		//	      "additionalProperties": false,
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"datastore_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomerManagedS3
				"customer_managed_s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Bucket
						"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(3, 255),
								stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9.\\-_]*"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: KeyPrefix
						"key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 255),
								stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9!_.*'()/{}:-]*/"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: RoleArn
						"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(20, 2048),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IotSiteWiseMultiLayerStorage
				"iot_site_wise_multi_layer_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CustomerManagedS3Storage
						"customer_managed_s3_storage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Bucket
								"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(3, 255),
										stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9.\\-_]*"), ""),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: KeyPrefix
								"key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 255),
										stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9!_.*'()/{}:-]*/"), ""),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ServiceManagedS3
				"service_managed_s3": schema.StringAttribute{ /*START ATTRIBUTE*/
					CustomType: jsontypes.NormalizedType{},
					Optional:   true,
					Computed:   true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileFormatConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "JsonConfiguration": {
		//	      "additionalProperties": false,
		//	      "type": "object"
		//	    },
		//	    "ParquetConfiguration": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "SchemaDefinition": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Columns": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Name": {
		//	                    "type": "string"
		//	                  },
		//	                  "Type": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "Type",
		//	                  "Name"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "maxItems": 100,
		//	              "minItems": 1,
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"file_format_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: JsonConfiguration
				"json_configuration": schema.StringAttribute{ /*START ATTRIBUTE*/
					CustomType: jsontypes.NormalizedType{},
					Optional:   true,
					Computed:   true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ParquetConfiguration
				"parquet_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: SchemaDefinition
						"schema_definition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Columns
								"columns": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													fwvalidators.NotNullString(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Type
											"type": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													fwvalidators.NotNullString(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Optional: true,
									Computed: true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeBetween(1, 100),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										generic.Multiset(),
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"datastore_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RetentionPeriod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "NumberOfDays": {
		//	      "maximum": 2147483647,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Unlimited": {
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"retention_period": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NumberOfDays
				"number_of_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 2147483647),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Unlimited
				"unlimited": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::IoTAnalytics::Datastore",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTAnalytics::Datastore").WithTerraformTypeName("awscc_iotanalytics_datastore")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attribute_name":                    "AttributeName",
		"bucket":                            "Bucket",
		"columns":                           "Columns",
		"customer_managed_s3":               "CustomerManagedS3",
		"customer_managed_s3_storage":       "CustomerManagedS3Storage",
		"datastore_id":                      "Id",
		"datastore_name":                    "DatastoreName",
		"datastore_partitions":              "DatastorePartitions",
		"datastore_storage":                 "DatastoreStorage",
		"file_format_configuration":         "FileFormatConfiguration",
		"iot_site_wise_multi_layer_storage": "IotSiteWiseMultiLayerStorage",
		"json_configuration":                "JsonConfiguration",
		"key":                               "Key",
		"key_prefix":                        "KeyPrefix",
		"name":                              "Name",
		"number_of_days":                    "NumberOfDays",
		"parquet_configuration":             "ParquetConfiguration",
		"partition":                         "Partition",
		"partitions":                        "Partitions",
		"retention_period":                  "RetentionPeriod",
		"role_arn":                          "RoleArn",
		"schema_definition":                 "SchemaDefinition",
		"service_managed_s3":                "ServiceManagedS3",
		"tags":                              "Tags",
		"timestamp_format":                  "TimestampFormat",
		"timestamp_partition":               "TimestampPartition",
		"type":                              "Type",
		"unlimited":                         "Unlimited",
		"value":                             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
