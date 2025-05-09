// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apprunner_service", serviceDataSource)
}

// serviceDataSource returns the Terraform awscc_apprunner_service data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppRunner::Service resource.
func serviceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoScalingConfigurationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Autoscaling configuration ARN",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"auto_scaling_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Autoscaling configuration ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Encryption configuration (KMS key)",
		//	  "properties": {
		//	    "KmsKey": {
		//	      "description": "The KMS Key",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "pattern": "arn:aws(-[\\w]+)*:kms:[a-z\\-]+-[0-9]{1}:[0-9]{12}:key\\/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "KmsKey"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKey
				"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The KMS Key",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Encryption configuration (KMS key)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Health check configuration",
		//	  "properties": {
		//	    "HealthyThreshold": {
		//	      "description": "Health check Healthy Threshold",
		//	      "maximum": 20,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Interval": {
		//	      "description": "Health check Interval",
		//	      "type": "integer"
		//	    },
		//	    "Path": {
		//	      "description": "Health check Path",
		//	      "type": "string"
		//	    },
		//	    "Protocol": {
		//	      "description": "Health Check Protocol",
		//	      "enum": [
		//	        "TCP",
		//	        "HTTP"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Timeout": {
		//	      "description": "Health check Timeout",
		//	      "maximum": 20,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "UnhealthyThreshold": {
		//	      "description": "Health check Unhealthy Threshold",
		//	      "maximum": 20,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"health_check_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: HealthyThreshold
				"healthy_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Health check Healthy Threshold",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Interval
				"interval": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Health check Interval",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Path
				"path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Health check Path",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Protocol
				"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Health Check Protocol",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Timeout
				"timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Health check Timeout",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UnhealthyThreshold
				"unhealthy_threshold": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Health check Unhealthy Threshold",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Health check configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Instance Configuration",
		//	  "properties": {
		//	    "Cpu": {
		//	      "description": "CPU",
		//	      "maxLength": 9,
		//	      "minLength": 3,
		//	      "pattern": "256|512|1024|2048|4096|(0.25|0.5|1|2|4) vCPU",
		//	      "type": "string"
		//	    },
		//	    "InstanceRoleArn": {
		//	      "description": "Instance Role Arn",
		//	      "maxLength": 1024,
		//	      "minLength": 29,
		//	      "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:role/[\\w+=,.@-]{1,64}",
		//	      "type": "string"
		//	    },
		//	    "Memory": {
		//	      "description": "Memory",
		//	      "maxLength": 6,
		//	      "minLength": 3,
		//	      "pattern": "512|1024|2048|3072|4096|6144|8192|10240|12288|(0.5|1|2|3|4|6|8|10|12) GB",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"instance_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Cpu
				"cpu": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "CPU",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: InstanceRoleArn
				"instance_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Instance Role Arn",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Memory
				"memory": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Memory",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Instance Configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Network configuration",
		//	  "properties": {
		//	    "EgressConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Network egress configuration",
		//	      "properties": {
		//	        "EgressType": {
		//	          "description": "Network egress type.",
		//	          "enum": [
		//	            "DEFAULT",
		//	            "VPC"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "VpcConnectorArn": {
		//	          "description": "The Amazon Resource Name (ARN) of the App Runner VpcConnector.",
		//	          "maxLength": 1011,
		//	          "minLength": 44,
		//	          "pattern": "",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "EgressType"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IngressConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Network ingress configuration",
		//	      "properties": {
		//	        "IsPubliclyAccessible": {
		//	          "description": "It's set to true if the Apprunner service is publicly accessible. It's set to false otherwise.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "IsPubliclyAccessible"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IpAddressType": {
		//	      "description": "App Runner service endpoint IP address type",
		//	      "enum": [
		//	        "IPV4",
		//	        "DUAL_STACK"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EgressConfiguration
				"egress_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EgressType
						"egress_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Network egress type.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: VpcConnectorArn
						"vpc_connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of the App Runner VpcConnector.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Network egress configuration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IngressConfiguration
				"ingress_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IsPubliclyAccessible
						"is_publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "It's set to true if the Apprunner service is publicly accessible. It's set to false otherwise.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Network ingress configuration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IpAddressType
				"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "App Runner service endpoint IP address type",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Network configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObservabilityConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Service observability configuration",
		//	  "properties": {
		//	    "ObservabilityConfigurationArn": {
		//	      "description": "The Amazon Resource Name (ARN) of the App Runner ObservabilityConfiguration.",
		//	      "maxLength": 1011,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "ObservabilityEnabled": {
		//	      "description": "Observability enabled",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "ObservabilityEnabled"
		//	  ],
		//	  "type": "object"
		//	}
		"observability_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ObservabilityConfigurationArn
				"observability_configuration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon Resource Name (ARN) of the App Runner ObservabilityConfiguration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ObservabilityEnabled
				"observability_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Observability enabled",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Service observability configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AppRunner Service.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AppRunner Service.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AppRunner Service Id",
		//	  "maxLength": 32,
		//	  "minLength": 32,
		//	  "type": "string"
		//	}
		"service_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AppRunner Service Id",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AppRunner Service Name.",
		//	  "maxLength": 40,
		//	  "minLength": 4,
		//	  "pattern": "[A-Za-z0-9][A-Za-z0-9-_]{3,39}",
		//	  "type": "string"
		//	}
		"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AppRunner Service Name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Service Url of the AppRunner Service.",
		//	  "type": "string"
		//	}
		"service_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Service Url of the AppRunner Service.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Source Code configuration",
		//	  "properties": {
		//	    "AuthenticationConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "Authentication Configuration",
		//	      "properties": {
		//	        "AccessRoleArn": {
		//	          "description": "Access Role Arn",
		//	          "maxLength": 1024,
		//	          "minLength": 29,
		//	          "pattern": "arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):iam::[0-9]{12}:role/[\\w+=,.@-]{1,64}",
		//	          "type": "string"
		//	        },
		//	        "ConnectionArn": {
		//	          "description": "Connection Arn",
		//	          "maxLength": 1011,
		//	          "minLength": 1,
		//	          "pattern": "",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "AutoDeploymentsEnabled": {
		//	      "description": "Auto Deployment enabled",
		//	      "type": "boolean"
		//	    },
		//	    "CodeRepository": {
		//	      "additionalProperties": false,
		//	      "description": "Source Code Repository",
		//	      "properties": {
		//	        "CodeConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Code Configuration",
		//	          "properties": {
		//	            "CodeConfigurationValues": {
		//	              "additionalProperties": false,
		//	              "description": "Code Configuration Values",
		//	              "properties": {
		//	                "BuildCommand": {
		//	                  "description": "Build Command",
		//	                  "type": "string"
		//	                },
		//	                "Port": {
		//	                  "description": "Port",
		//	                  "type": "string"
		//	                },
		//	                "Runtime": {
		//	                  "description": "Runtime",
		//	                  "enum": [
		//	                    "PYTHON_3",
		//	                    "NODEJS_12",
		//	                    "NODEJS_14",
		//	                    "CORRETTO_8",
		//	                    "CORRETTO_11",
		//	                    "NODEJS_16",
		//	                    "GO_1",
		//	                    "DOTNET_6",
		//	                    "PHP_81",
		//	                    "RUBY_31",
		//	                    "PYTHON_311",
		//	                    "NODEJS_18",
		//	                    "NODEJS_22"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "RuntimeEnvironmentSecrets": {
		//	                  "description": "The secrets and parameters that get referenced by your service as environment variables",
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "Name": {
		//	                        "type": "string"
		//	                      },
		//	                      "Value": {
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "RuntimeEnvironmentVariables": {
		//	                  "items": {
		//	                    "additionalProperties": false,
		//	                    "properties": {
		//	                      "Name": {
		//	                        "type": "string"
		//	                      },
		//	                      "Value": {
		//	                        "type": "string"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "StartCommand": {
		//	                  "description": "Start Command",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Runtime"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "ConfigurationSource": {
		//	              "description": "Configuration Source",
		//	              "enum": [
		//	                "REPOSITORY",
		//	                "API"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ConfigurationSource"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "RepositoryUrl": {
		//	          "description": "Repository Url",
		//	          "type": "string"
		//	        },
		//	        "SourceCodeVersion": {
		//	          "additionalProperties": false,
		//	          "description": "Source Code Version",
		//	          "properties": {
		//	            "Type": {
		//	              "description": "Source Code Version Type",
		//	              "enum": [
		//	                "BRANCH"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "description": "Source Code Version Value",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Type",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "SourceDirectory": {
		//	          "description": "Source Directory",
		//	          "maxLength": 4096,
		//	          "minLength": 1,
		//	          "pattern": "[^\\x00]+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "RepositoryUrl",
		//	        "SourceCodeVersion"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "ImageRepository": {
		//	      "additionalProperties": false,
		//	      "description": "Image Repository",
		//	      "properties": {
		//	        "ImageConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Image Configuration",
		//	          "properties": {
		//	            "Port": {
		//	              "description": "Port",
		//	              "type": "string"
		//	            },
		//	            "RuntimeEnvironmentSecrets": {
		//	              "description": "The secrets and parameters that get referenced by your service as environment variables",
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Name": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "RuntimeEnvironmentVariables": {
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Name": {
		//	                    "type": "string"
		//	                  },
		//	                  "Value": {
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "type": "array"
		//	            },
		//	            "StartCommand": {
		//	              "description": "Start Command",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ImageIdentifier": {
		//	          "description": "Image Identifier",
		//	          "maxLength": 1024,
		//	          "minLength": 1,
		//	          "pattern": "([0-9]{12}.dkr.ecr.[a-z\\-]+-[0-9]{1}.amazonaws.com\\/.*)|(^public\\.ecr\\.aws\\/.+\\/.+)",
		//	          "type": "string"
		//	        },
		//	        "ImageRepositoryType": {
		//	          "description": "Image Repository Type",
		//	          "enum": [
		//	            "ECR",
		//	            "ECR_PUBLIC"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ImageIdentifier",
		//	        "ImageRepositoryType"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthenticationConfiguration
				"authentication_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AccessRoleArn
						"access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Access Role Arn",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ConnectionArn
						"connection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Connection Arn",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Authentication Configuration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: AutoDeploymentsEnabled
				"auto_deployments_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Auto Deployment enabled",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CodeRepository
				"code_repository": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CodeConfiguration
						"code_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CodeConfigurationValues
								"code_configuration_values": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: BuildCommand
										"build_command": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Build Command",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Port
										"port": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Port",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Runtime
										"runtime": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Runtime",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: RuntimeEnvironmentSecrets
										"runtime_environment_secrets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Name
													"name": schema.StringAttribute{ /*START ATTRIBUTE*/
														Computed: true,
													}, /*END ATTRIBUTE*/
													// Property: Value
													"value": schema.StringAttribute{ /*START ATTRIBUTE*/
														Computed: true,
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Description: "The secrets and parameters that get referenced by your service as environment variables",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: RuntimeEnvironmentVariables
										"runtime_environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Name
													"name": schema.StringAttribute{ /*START ATTRIBUTE*/
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
										// Property: StartCommand
										"start_command": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Start Command",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Description: "Code Configuration Values",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: ConfigurationSource
								"configuration_source": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Configuration Source",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Code Configuration",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: RepositoryUrl
						"repository_url": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Repository Url",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: SourceCodeVersion
						"source_code_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Type
								"type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Source Code Version Type",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Source Code Version Value",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Source Code Version",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: SourceDirectory
						"source_directory": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Source Directory",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Source Code Repository",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ImageRepository
				"image_repository": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ImageConfiguration
						"image_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Port
								"port": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Port",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: RuntimeEnvironmentSecrets
								"runtime_environment_secrets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Value
											"value": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "The secrets and parameters that get referenced by your service as environment variables",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: RuntimeEnvironmentVariables
								"runtime_environment_variables": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
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
								// Property: StartCommand
								"start_command": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Start Command",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Image Configuration",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ImageIdentifier
						"image_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Image Identifier",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ImageRepositoryType
						"image_repository_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Image Repository Type",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Image Repository",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Source Code configuration",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AppRunner Service status.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AppRunner Service status.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
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
		Description: "Data Source schema for AWS::AppRunner::Service",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::Service").WithTerraformTypeName("awscc_apprunner_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_role_arn":                 "AccessRoleArn",
		"authentication_configuration":    "AuthenticationConfiguration",
		"auto_deployments_enabled":        "AutoDeploymentsEnabled",
		"auto_scaling_configuration_arn":  "AutoScalingConfigurationArn",
		"build_command":                   "BuildCommand",
		"code_configuration":              "CodeConfiguration",
		"code_configuration_values":       "CodeConfigurationValues",
		"code_repository":                 "CodeRepository",
		"configuration_source":            "ConfigurationSource",
		"connection_arn":                  "ConnectionArn",
		"cpu":                             "Cpu",
		"egress_configuration":            "EgressConfiguration",
		"egress_type":                     "EgressType",
		"encryption_configuration":        "EncryptionConfiguration",
		"health_check_configuration":      "HealthCheckConfiguration",
		"healthy_threshold":               "HealthyThreshold",
		"image_configuration":             "ImageConfiguration",
		"image_identifier":                "ImageIdentifier",
		"image_repository":                "ImageRepository",
		"image_repository_type":           "ImageRepositoryType",
		"ingress_configuration":           "IngressConfiguration",
		"instance_configuration":          "InstanceConfiguration",
		"instance_role_arn":               "InstanceRoleArn",
		"interval":                        "Interval",
		"ip_address_type":                 "IpAddressType",
		"is_publicly_accessible":          "IsPubliclyAccessible",
		"key":                             "Key",
		"kms_key":                         "KmsKey",
		"memory":                          "Memory",
		"name":                            "Name",
		"network_configuration":           "NetworkConfiguration",
		"observability_configuration":     "ObservabilityConfiguration",
		"observability_configuration_arn": "ObservabilityConfigurationArn",
		"observability_enabled":           "ObservabilityEnabled",
		"path":                            "Path",
		"port":                            "Port",
		"protocol":                        "Protocol",
		"repository_url":                  "RepositoryUrl",
		"runtime":                         "Runtime",
		"runtime_environment_secrets":     "RuntimeEnvironmentSecrets",
		"runtime_environment_variables":   "RuntimeEnvironmentVariables",
		"service_arn":                     "ServiceArn",
		"service_id":                      "ServiceId",
		"service_name":                    "ServiceName",
		"service_url":                     "ServiceUrl",
		"source_code_version":             "SourceCodeVersion",
		"source_configuration":            "SourceConfiguration",
		"source_directory":                "SourceDirectory",
		"start_command":                   "StartCommand",
		"status":                          "Status",
		"tags":                            "Tags",
		"timeout":                         "Timeout",
		"type":                            "Type",
		"unhealthy_threshold":             "UnhealthyThreshold",
		"value":                           "Value",
		"vpc_connector_arn":               "VpcConnectorArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
