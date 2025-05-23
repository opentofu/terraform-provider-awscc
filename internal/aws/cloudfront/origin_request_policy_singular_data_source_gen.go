// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_origin_request_policy", originRequestPolicyDataSource)
}

// originRequestPolicyDataSource returns the Terraform awscc_cloudfront_origin_request_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::OriginRequestPolicy resource.
func originRequestPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"origin_request_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OriginRequestPolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The origin request policy configuration.",
		//	  "properties": {
		//	    "Comment": {
		//	      "description": "A comment to describe the origin request policy. The comment cannot be longer than 128 characters.",
		//	      "type": "string"
		//	    },
		//	    "CookiesConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The cookies from viewer requests to include in origin requests.",
		//	      "properties": {
		//	        "CookieBehavior": {
		//	          "description": "Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No cookies in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any cookies that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the cookies in viewer requests that are listed in the ``CookieNames`` type are included in requests that CloudFront sends to the origin.\n  +  ``all`` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``CookieNames`` type, which are not included.",
		//	          "pattern": "^(none|whitelist|all|allExcept)$",
		//	          "type": "string"
		//	        },
		//	        "Cookies": {
		//	          "description": "Contains a list of cookie names.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "CookieBehavior"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "HeadersConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The HTTP headers to include in origin requests. These can include headers from viewer requests and additional headers added by CloudFront.",
		//	      "properties": {
		//	        "HeaderBehavior": {
		//	          "description": "Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No HTTP headers in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any headers that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the HTTP headers that are listed in the ``Headers`` type are included in requests that CloudFront sends to the origin.\n  +  ``allViewer`` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allViewerAndWhitelistCloudFront`` ? All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the ``Headers`` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.\n  +  ``allExcept`` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``Headers`` type, which are not included.",
		//	          "pattern": "^(none|whitelist|allViewer|allViewerAndWhitelistCloudFront|allExcept)$",
		//	          "type": "string"
		//	        },
		//	        "Headers": {
		//	          "description": "Contains a list of HTTP header names.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "HeaderBehavior"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Name": {
		//	      "description": "A unique name to identify the origin request policy.",
		//	      "type": "string"
		//	    },
		//	    "QueryStringsConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The URL query strings from viewer requests to include in origin requests.",
		//	      "properties": {
		//	        "QueryStringBehavior": {
		//	          "description": "Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in requests that CloudFront sends to the origin.\n  +  ``all`` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``QueryStringNames`` type, which are not included.",
		//	          "pattern": "^(none|whitelist|all|allExcept)$",
		//	          "type": "string"
		//	        },
		//	        "QueryStrings": {
		//	          "description": "Contains a list of query string names.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        }
		//	      },
		//	      "required": [
		//	        "QueryStringBehavior"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "HeadersConfig",
		//	    "CookiesConfig",
		//	    "QueryStringsConfig"
		//	  ],
		//	  "type": "object"
		//	}
		"origin_request_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comment to describe the origin request policy. The comment cannot be longer than 128 characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CookiesConfig
				"cookies_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CookieBehavior
						"cookie_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Determines whether cookies in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No cookies in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any cookies that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the cookies in viewer requests that are listed in the ``CookieNames`` type are included in requests that CloudFront sends to the origin.\n  +  ``all`` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All cookies in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``CookieNames`` type, which are not included.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Cookies
						"cookies": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Contains a list of cookie names.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The cookies from viewer requests to include in origin requests.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: HeadersConfig
				"headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: HeaderBehavior
						"header_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Determines whether any HTTP headers are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No HTTP headers in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any headers that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the HTTP headers that are listed in the ``Headers`` type are included in requests that CloudFront sends to the origin.\n  +  ``allViewer`` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allViewerAndWhitelistCloudFront`` ? All HTTP headers in viewer requests and the additional CloudFront headers that are listed in the ``Headers`` type are included in requests that CloudFront sends to the origin. The additional headers are added by CloudFront.\n  +  ``allExcept`` ? All HTTP headers in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``Headers`` type, which are not included.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Headers
						"headers": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Contains a list of HTTP header names.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The HTTP headers to include in origin requests. These can include headers from viewer requests and additional headers added by CloudFront.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A unique name to identify the origin request policy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: QueryStringsConfig
				"query_strings_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: QueryStringBehavior
						"query_string_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in a ``CachePolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in requests that CloudFront sends to the origin.\n  +  ``all`` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All query strings in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``QueryStringNames`` type, which are not included.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: QueryStrings
						"query_strings": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Contains a list of query string names.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The URL query strings from viewer requests to include in origin requests.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The origin request policy configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::OriginRequestPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::OriginRequestPolicy").WithTerraformTypeName("awscc_cloudfront_origin_request_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"comment":                      "Comment",
		"cookie_behavior":              "CookieBehavior",
		"cookies":                      "Cookies",
		"cookies_config":               "CookiesConfig",
		"header_behavior":              "HeaderBehavior",
		"headers":                      "Headers",
		"headers_config":               "HeadersConfig",
		"last_modified_time":           "LastModifiedTime",
		"name":                         "Name",
		"origin_request_policy_config": "OriginRequestPolicyConfig",
		"origin_request_policy_id":     "Id",
		"query_string_behavior":        "QueryStringBehavior",
		"query_strings":                "QueryStrings",
		"query_strings_config":         "QueryStringsConfig",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
