// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cleanrooms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cleanrooms_configured_table", configuredTableDataSource)
}

// configuredTableDataSource returns the Terraform awscc_cleanrooms_configured_table data source.
// This Terraform data source corresponds to the CloudFormation AWS::CleanRooms::ConfiguredTable resource.
func configuredTableDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllowedColumns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 100,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"allowed_columns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AnalysisMethod
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DIRECT_QUERY"
		//	  ],
		//	  "type": "string"
		//	}
		"analysis_method": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AnalysisRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Policy": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "V1": {
		//	            "properties": {
		//	              "Aggregation": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "AggregateColumns": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "additionalProperties": false,
		//	                      "properties": {
		//	                        "ColumnNames": {
		//	                          "insertionOrder": false,
		//	                          "items": {
		//	                            "maxLength": 127,
		//	                            "minLength": 1,
		//	                            "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                            "type": "string"
		//	                          },
		//	                          "minItems": 1,
		//	                          "type": "array"
		//	                        },
		//	                        "Function": {
		//	                          "enum": [
		//	                            "SUM",
		//	                            "SUM_DISTINCT",
		//	                            "COUNT",
		//	                            "COUNT_DISTINCT",
		//	                            "AVG"
		//	                          ],
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "ColumnNames",
		//	                        "Function"
		//	                      ],
		//	                      "type": "object"
		//	                    },
		//	                    "minItems": 1,
		//	                    "type": "array"
		//	                  },
		//	                  "AllowedJoinOperators": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "enum": [
		//	                        "OR",
		//	                        "AND"
		//	                      ],
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 2,
		//	                    "type": "array"
		//	                  },
		//	                  "DimensionColumns": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 127,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                      "type": "string"
		//	                    },
		//	                    "type": "array"
		//	                  },
		//	                  "JoinColumns": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 127,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                      "type": "string"
		//	                    },
		//	                    "type": "array"
		//	                  },
		//	                  "JoinRequired": {
		//	                    "enum": [
		//	                      "QUERY_RUNNER"
		//	                    ],
		//	                    "type": "string"
		//	                  },
		//	                  "OutputConstraints": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "additionalProperties": false,
		//	                      "properties": {
		//	                        "ColumnName": {
		//	                          "maxLength": 127,
		//	                          "minLength": 1,
		//	                          "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                          "type": "string"
		//	                        },
		//	                        "Minimum": {
		//	                          "maximum": 100000,
		//	                          "minimum": 2,
		//	                          "type": "number"
		//	                        },
		//	                        "Type": {
		//	                          "enum": [
		//	                            "COUNT_DISTINCT"
		//	                          ],
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "ColumnName",
		//	                        "Minimum",
		//	                        "Type"
		//	                      ],
		//	                      "type": "object"
		//	                    },
		//	                    "minItems": 1,
		//	                    "type": "array"
		//	                  },
		//	                  "ScalarFunctions": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "enum": [
		//	                        "TRUNC",
		//	                        "ABS",
		//	                        "CEILING",
		//	                        "FLOOR",
		//	                        "LN",
		//	                        "LOG",
		//	                        "ROUND",
		//	                        "SQRT",
		//	                        "CAST",
		//	                        "LOWER",
		//	                        "RTRIM",
		//	                        "UPPER",
		//	                        "COALESCE"
		//	                      ],
		//	                      "type": "string"
		//	                    },
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "AggregateColumns",
		//	                  "JoinColumns",
		//	                  "DimensionColumns",
		//	                  "ScalarFunctions",
		//	                  "OutputConstraints"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "List": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "AllowedJoinOperators": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "enum": [
		//	                        "OR",
		//	                        "AND"
		//	                      ],
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 2,
		//	                    "type": "array"
		//	                  },
		//	                  "JoinColumns": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 127,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                      "type": "string"
		//	                    },
		//	                    "minItems": 1,
		//	                    "type": "array"
		//	                  },
		//	                  "ListColumns": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 127,
		//	                      "minLength": 1,
		//	                      "pattern": "^[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?$",
		//	                      "type": "string"
		//	                    },
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "JoinColumns",
		//	                  "ListColumns"
		//	                ],
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "V1"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "AGGREGATION",
		//	          "LIST"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Type",
		//	      "Policy"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"analysis_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Policy
					"policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: V1
							"v1": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Aggregation
									"aggregation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: AggregateColumns
											"aggregate_columns": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
												NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: ColumnNames
														"column_names": schema.ListAttribute{ /*START ATTRIBUTE*/
															ElementType: types.StringType,
															Computed:    true,
														}, /*END ATTRIBUTE*/
														// Property: Function
														"function": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
												}, /*END NESTED OBJECT*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: AllowedJoinOperators
											"allowed_join_operators": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: DimensionColumns
											"dimension_columns": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: JoinColumns
											"join_columns": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: JoinRequired
											"join_required": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: OutputConstraints
											"output_constraints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
												NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: ColumnName
														"column_name": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
														// Property: Minimum
														"minimum": schema.Float64Attribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
														// Property: Type
														"type": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
												}, /*END NESTED OBJECT*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: ScalarFunctions
											"scalar_functions": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: List
									"list": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: AllowedJoinOperators
											"allowed_join_operators": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: JoinColumns
											"join_columns": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: ListColumns
											"list_columns": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ConfiguredTableIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"configured_table_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TableReference
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Glue": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DatabaseName": {
		//	          "maxLength": 128,
		//	          "pattern": "^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$",
		//	          "type": "string"
		//	        },
		//	        "TableName": {
		//	          "maxLength": 128,
		//	          "pattern": "^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "DatabaseName",
		//	        "TableName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Glue"
		//	  ],
		//	  "type": "object"
		//	}
		"table_reference": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Glue
				"glue": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DatabaseName
						"database_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TableName
						"table_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CleanRooms::ConfiguredTable",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CleanRooms::ConfiguredTable").WithTerraformTypeName("awscc_cleanrooms_configured_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregate_columns":           "AggregateColumns",
		"aggregation":                 "Aggregation",
		"allowed_columns":             "AllowedColumns",
		"allowed_join_operators":      "AllowedJoinOperators",
		"analysis_method":             "AnalysisMethod",
		"analysis_rules":              "AnalysisRules",
		"arn":                         "Arn",
		"column_name":                 "ColumnName",
		"column_names":                "ColumnNames",
		"configured_table_identifier": "ConfiguredTableIdentifier",
		"database_name":               "DatabaseName",
		"description":                 "Description",
		"dimension_columns":           "DimensionColumns",
		"function":                    "Function",
		"glue":                        "Glue",
		"join_columns":                "JoinColumns",
		"join_required":               "JoinRequired",
		"key":                         "Key",
		"list":                        "List",
		"list_columns":                "ListColumns",
		"minimum":                     "Minimum",
		"name":                        "Name",
		"output_constraints":          "OutputConstraints",
		"policy":                      "Policy",
		"scalar_functions":            "ScalarFunctions",
		"table_name":                  "TableName",
		"table_reference":             "TableReference",
		"tags":                        "Tags",
		"type":                        "Type",
		"v1":                          "V1",
		"value":                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}