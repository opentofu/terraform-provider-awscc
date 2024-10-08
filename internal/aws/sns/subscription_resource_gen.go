// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_sns_subscription", subscriptionResource)
}

// subscriptionResource returns the Terraform awscc_sns_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::SNS::Subscription resource.
func subscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn of the subscription",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn of the subscription",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.",
		//	  "type": "string"
		//	}
		"delivery_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The delivery policy JSON assigned to the subscription. Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subscription's endpoint. The endpoint value depends on the protocol that you specify. ",
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subscription's endpoint. The endpoint value depends on the protocol that you specify. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FilterPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.",
		//	  "type": "string"
		//	}
		"filter_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The filter policy JSON assigned to the subscription. Enables the subscriber to filter out unwanted messages.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FilterPolicyScope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.",
		//	  "type": "string"
		//	}
		"filter_policy_scope": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This attribute lets you choose the filtering scope by using one of the following string value types: MessageAttributes (default) and MessageBody.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Protocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subscription's protocol.",
		//	  "type": "string"
		//	}
		"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subscription's protocol.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RawMessageDelivery
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.",
		//	  "type": "boolean"
		//	}
		"raw_message_delivery": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When set to true, enables raw message delivery. Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RedrivePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.",
		//	  "type": "string"
		//	}
		"redrive_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue. Messages that can't be delivered due to client errors are held in the dead-letter queue for further analysis or reprocessing.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Region
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.",
		//	  "type": "string"
		//	}
		"region": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For cross-region subscriptions, the region in which the topic resides.If no region is specified, AWS CloudFormation uses the region of the caller as the default.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Region is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ReplayPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.",
		//	  "type": "string"
		//	}
		"replay_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether Amazon SNS resends the notification to the subscription when a message's attribute changes.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This property applies only to Amazon Data Firehose delivery stream subscriptions.",
		//	  "type": "string"
		//	}
		"subscription_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This property applies only to Amazon Data Firehose delivery stream subscriptions.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TopicArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the topic to subscribe to.",
		//	  "type": "string"
		//	}
		"topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the topic to subscribe to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::SNS::Subscription",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SNS::Subscription").WithTerraformTypeName("awscc_sns_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"delivery_policy":       "DeliveryPolicy",
		"endpoint":              "Endpoint",
		"filter_policy":         "FilterPolicy",
		"filter_policy_scope":   "FilterPolicyScope",
		"protocol":              "Protocol",
		"raw_message_delivery":  "RawMessageDelivery",
		"redrive_policy":        "RedrivePolicy",
		"region":                "Region",
		"replay_policy":         "ReplayPolicy",
		"subscription_role_arn": "SubscriptionRoleArn",
		"topic_arn":             "TopicArn",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Region",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
