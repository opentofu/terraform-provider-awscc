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
	registry.AddDataSourceFactory("awscc_cloudfront_cache_policy", cachePolicyDataSource)
}

// cachePolicyDataSource returns the Terraform awscc_cloudfront_cache_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::CachePolicy resource.
func cachePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CachePolicyConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The cache policy configuration.",
		//	  "properties": {
		//	    "Comment": {
		//	      "description": "A comment to describe the cache policy. The comment cannot be longer than 128 characters.",
		//	      "type": "string"
		//	    },
		//	    "DefaultTTL": {
		//	      "description": "The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 86400 seconds (one day). If the value of ``MinTTL`` is more than 86400 seconds, then the default value for this field is the same as the value of ``MinTTL``.",
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "MaxTTL": {
		//	      "description": "The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value only when the origin sends ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 31536000 seconds (one year). If the value of ``MinTTL`` or ``DefaultTTL`` is more than 31536000 seconds, then the default value for this field is the same as the value of ``DefaultTTL``.",
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "MinTTL": {
		//	      "description": "The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.",
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "Name": {
		//	      "description": "A unique name to identify the cache policy.",
		//	      "type": "string"
		//	    },
		//	    "ParametersInCacheKeyAndForwardedToOrigin": {
		//	      "additionalProperties": false,
		//	      "description": "The HTTP headers, cookies, and URL query strings to include in the cache key. The values included in the cache key are also included in requests that CloudFront sends to the origin.",
		//	      "properties": {
		//	        "CookiesConfig": {
		//	          "additionalProperties": false,
		//	          "description": "An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and in requests that CloudFront sends to the origin.",
		//	          "properties": {
		//	            "CookieBehavior": {
		//	              "description": "Determines whether any cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No cookies in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any cookies that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the cookies in viewer requests that are listed in the ``CookieNames`` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* for those that are listed in the ``CookieNames`` type, which are not included.\n  +  ``all`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
		//	              "pattern": "^(none|whitelist|allExcept|all)$",
		//	              "type": "string"
		//	            },
		//	            "Cookies": {
		//	              "description": "Contains a list of cookie names.",
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "CookieBehavior"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "EnableAcceptEncodingBrotli": {
		//	          "description": "A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ``EnableAcceptEncodingGzip`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ``Accept-Encoding`` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.\n If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.",
		//	          "type": "boolean"
		//	        },
		//	        "EnableAcceptEncodingGzip": {
		//	          "description": "A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ``EnableAcceptEncodingBrotli`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ``Accept-Encoding`` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.\n If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.",
		//	          "type": "boolean"
		//	        },
		//	        "HeadersConfig": {
		//	          "additionalProperties": false,
		//	          "description": "An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and in requests that CloudFront sends to the origin.",
		//	          "properties": {
		//	            "HeaderBehavior": {
		//	              "description": "Determines whether any HTTP headers are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No HTTP headers are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any headers that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the HTTP headers that are listed in the ``Headers`` type are included in the cache key and in requests that CloudFront sends to the origin.",
		//	              "pattern": "^(none|whitelist)$",
		//	              "type": "string"
		//	            },
		//	            "Headers": {
		//	              "description": "Contains a list of HTTP header names.",
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "HeaderBehavior"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "QueryStringsConfig": {
		//	          "additionalProperties": false,
		//	          "description": "An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and in requests that CloudFront sends to the origin.",
		//	          "properties": {
		//	            "QueryStringBehavior": {
		//	              "description": "Determines whether any URL query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No query strings in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* those that are listed in the ``QueryStringNames`` type, which are not included.\n  +  ``all`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
		//	              "pattern": "^(none|whitelist|allExcept|all)$",
		//	              "type": "string"
		//	            },
		//	            "QueryStrings": {
		//	              "description": "Contains a list of query string names.",
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "required": [
		//	            "QueryStringBehavior"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "EnableAcceptEncodingGzip",
		//	        "HeadersConfig",
		//	        "CookiesConfig",
		//	        "QueryStringsConfig"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Name",
		//	    "MinTTL",
		//	    "MaxTTL",
		//	    "DefaultTTL",
		//	    "ParametersInCacheKeyAndForwardedToOrigin"
		//	  ],
		//	  "type": "object"
		//	}
		"cache_policy_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Comment
				"comment": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comment to describe the cache policy. The comment cannot be longer than 128 characters.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DefaultTTL
				"default_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 86400 seconds (one day). If the value of ``MinTTL`` is more than 86400 seconds, then the default value for this field is the same as the value of ``MinTTL``.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxTTL
				"max_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value only when the origin sends ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.\n The default value for this field is 31536000 seconds (one year). If the value of ``MinTTL`` or ``DefaultTTL`` is more than 31536000 seconds, then the default value for this field is the same as the value of ``DefaultTTL``.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MinTTL
				"min_ttl": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A unique name to identify the cache policy.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ParametersInCacheKeyAndForwardedToOrigin
				"parameters_in_cache_key_and_forwarded_to_origin": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CookiesConfig
						"cookies_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CookieBehavior
								"cookie_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Determines whether any cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No cookies in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any cookies that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the cookies in viewer requests that are listed in the ``CookieNames`` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* for those that are listed in the ``CookieNames`` type, which are not included.\n  +  ``all`` ? All cookies in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Cookies
								"cookies": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "Contains a list of cookie names.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and in requests that CloudFront sends to the origin.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: EnableAcceptEncodingBrotli
						"enable_accept_encoding_brotli": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ``EnableAcceptEncodingGzip`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ``Accept-Encoding`` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.\n If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: EnableAcceptEncodingGzip
						"enable_accept_encoding_gzip": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "A flag that can affect whether the ``Accept-Encoding`` HTTP header is included in the cache key and included in requests that CloudFront sends to the origin.\n This field is related to the ``EnableAcceptEncodingBrotli`` field. If one or both of these fields is ``true``*and* the viewer request includes the ``Accept-Encoding`` header, then CloudFront does the following:\n  +  Normalizes the value of the viewer's ``Accept-Encoding`` header\n  +  Includes the normalized header in the cache key\n  +  Includes the normalized header in the request to the origin, if a request is necessary\n  \n For more information, see [Compression support](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-policy-compressed-objects) in the *Amazon CloudFront Developer Guide*.\n If you set this value to ``true``, and this cache behavior also has an origin request policy attached, do not include the ``Accept-Encoding`` header in the origin request policy. CloudFront always includes the ``Accept-Encoding`` header in origin requests when the value of this field is ``true``, so including this header in an origin request policy has no effect.\n If both of these fields are ``false``, then CloudFront treats the ``Accept-Encoding`` header the same as any other HTTP header in the viewer request. By default, it's not included in the cache key and it's not included in origin requests. In this case, you can manually add ``Accept-Encoding`` to the headers whitelist like any other HTTP header.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: HeadersConfig
						"headers_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: HeaderBehavior
								"header_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Determines whether any HTTP headers are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No HTTP headers are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any headers that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the HTTP headers that are listed in the ``Headers`` type are included in the cache key and in requests that CloudFront sends to the origin.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: Headers
								"headers": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "Contains a list of HTTP header names.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "An object that determines whether any HTTP headers (and if so, which headers) are included in the cache key and in requests that CloudFront sends to the origin.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: QueryStringsConfig
						"query_strings_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: QueryStringBehavior
								"query_string_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Determines whether any URL query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin. Valid values are:\n  +  ``none`` ? No query strings in viewer requests are included in the cache key or in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in an ``OriginRequestPolicy``*are* included in origin requests.\n  +  ``whitelist`` ? Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in the cache key and in requests that CloudFront sends to the origin.\n  +  ``allExcept`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin, *except* those that are listed in the ``QueryStringNames`` type, which are not included.\n  +  ``all`` ? All query strings in viewer requests are included in the cache key and in requests that CloudFront sends to the origin.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: QueryStrings
								"query_strings": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "Contains a list of query string names.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the cache key and in requests that CloudFront sends to the origin.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The HTTP headers, cookies, and URL query strings to include in the cache key. The values included in the cache key are also included in requests that CloudFront sends to the origin.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The cache policy configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"cache_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFront::CachePolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::CachePolicy").WithTerraformTypeName("awscc_cloudfront_cache_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cache_policy_config":           "CachePolicyConfig",
		"cache_policy_id":               "Id",
		"comment":                       "Comment",
		"cookie_behavior":               "CookieBehavior",
		"cookies":                       "Cookies",
		"cookies_config":                "CookiesConfig",
		"default_ttl":                   "DefaultTTL",
		"enable_accept_encoding_brotli": "EnableAcceptEncodingBrotli",
		"enable_accept_encoding_gzip":   "EnableAcceptEncodingGzip",
		"header_behavior":               "HeaderBehavior",
		"headers":                       "Headers",
		"headers_config":                "HeadersConfig",
		"last_modified_time":            "LastModifiedTime",
		"max_ttl":                       "MaxTTL",
		"min_ttl":                       "MinTTL",
		"name":                          "Name",
		"parameters_in_cache_key_and_forwarded_to_origin": "ParametersInCacheKeyAndForwardedToOrigin",
		"query_string_behavior":                           "QueryStringBehavior",
		"query_strings":                                   "QueryStrings",
		"query_strings_config":                            "QueryStringsConfig",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
