// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package applicationautoscaling

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_applicationautoscaling_scalable_target", scalableTargetResource)
}

// scalableTargetResource returns the Terraform awscc_applicationautoscaling_scalable_target resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApplicationAutoScaling::ScalableTarget resource.
func scalableTargetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"scalable_target_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum value that you plan to scale out to. When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.",
		//	  "type": "integer"
		//	}
		"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum value that you plan to scale out to. When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: MinCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.",
		//	  "type": "integer"
		//	}
		"min_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "anyOf": [
		//	    {},
		//	    {}
		//	  ],
		//	  "description": "The identifier of the resource associated with the scalable target. This string consists of the resource type and unique identifier.\n  +  ECS service - The resource type is ``service`` and the unique identifier is the cluster name and service name. Example: ``service/my-cluster/my-service``.\n  +  Spot Fleet - The resource type is ``spot-fleet-request`` and the unique identifier is the Spot Fleet request ID. Example: ``spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE``.\n  +  EMR cluster - The resource type is ``instancegroup`` and the unique identifier is the cluster ID and instance group ID. Example: ``instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0``.\n  +  AppStream 2.0 fleet - The resource type is ``fleet`` and the unique identifier is the fleet name. Example: ``fleet/sample-fleet``.\n  +  DynamoDB table - The resource type is ``table`` and the unique identifier is the table name. Example: ``table/my-table``.\n  +  DynamoDB global secondary index - The resource type is ``index`` and the unique identifier is the index name. Example: ``table/my-table/index/my-table-index``.\n  +  Aurora DB cluster - The resource type is ``cluster`` and the unique identifier is the cluster name. Example: ``cluster:my-db-cluster``.\n  +  SageMaker endpoint variant - The resource type is ``variant`` and the unique identifier is the resource ID. Example: ``endpoint/my-end-point/variant/KMeansClustering``.\n  +  Custom resources are not supported with a resource type. This parameter must specify the ``OutputValue`` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource).\n  +  Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ``arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE``.\n  +  Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ``arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE``.\n  +  Lambda provisioned concurrency - The resource type is ``function`` and the unique identifier is the function name with a function version or alias name suffix that is not ``$LATEST``. Example: ``function:my-function:prod`` or ``function:my-function:1``.\n  +  Amazon Keyspaces table - The resource type is ``table`` and the unique identifier is the table name. Example: ``keyspace/mykeyspace/table/mytable``.\n  +  Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: ``arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5``.\n  +  Amazon ElastiCache replication group - The resource type is ``replication-group`` and the unique identifier is the replication group name. Example: ``replication-group/mycluster``.\n  +  Neptune cluster - The resource type is ``cluster`` and the unique identifier is the cluster name. Example: ``cluster:mycluster``.\n  +  SageMaker serverless endpoint - The resource type is ``variant`` and the unique identifier is the resource ID. Example: ``endpoint/my-end-point/variant/KMeansClustering``.\n  +  SageMaker inference component - The resource type is ``inference-component`` and the unique identifier is the resource ID. Example: ``inference-component/my-inference-component``.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the resource associated with the scalable target. This string consists of the resource type and unique identifier.\n  +  ECS service - The resource type is ``service`` and the unique identifier is the cluster name and service name. Example: ``service/my-cluster/my-service``.\n  +  Spot Fleet - The resource type is ``spot-fleet-request`` and the unique identifier is the Spot Fleet request ID. Example: ``spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE``.\n  +  EMR cluster - The resource type is ``instancegroup`` and the unique identifier is the cluster ID and instance group ID. Example: ``instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0``.\n  +  AppStream 2.0 fleet - The resource type is ``fleet`` and the unique identifier is the fleet name. Example: ``fleet/sample-fleet``.\n  +  DynamoDB table - The resource type is ``table`` and the unique identifier is the table name. Example: ``table/my-table``.\n  +  DynamoDB global secondary index - The resource type is ``index`` and the unique identifier is the index name. Example: ``table/my-table/index/my-table-index``.\n  +  Aurora DB cluster - The resource type is ``cluster`` and the unique identifier is the cluster name. Example: ``cluster:my-db-cluster``.\n  +  SageMaker endpoint variant - The resource type is ``variant`` and the unique identifier is the resource ID. Example: ``endpoint/my-end-point/variant/KMeansClustering``.\n  +  Custom resources are not supported with a resource type. This parameter must specify the ``OutputValue`` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource).\n  +  Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ``arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE``.\n  +  Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: ``arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE``.\n  +  Lambda provisioned concurrency - The resource type is ``function`` and the unique identifier is the function name with a function version or alias name suffix that is not ``$LATEST``. Example: ``function:my-function:prod`` or ``function:my-function:1``.\n  +  Amazon Keyspaces table - The resource type is ``table`` and the unique identifier is the table name. Example: ``keyspace/mykeyspace/table/mytable``.\n  +  Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: ``arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5``.\n  +  Amazon ElastiCache replication group - The resource type is ``replication-group`` and the unique identifier is the replication group name. Example: ``replication-group/mycluster``.\n  +  Neptune cluster - The resource type is ``cluster`` and the unique identifier is the cluster name. Example: ``cluster:mycluster``.\n  +  SageMaker serverless endpoint - The resource type is ``variant`` and the unique identifier is the resource ID. Example: ``endpoint/my-end-point/variant/KMeansClustering``.\n  +  SageMaker inference component - The resource type is ``inference-component`` and the unique identifier is the resource ID. Example: ``inference-component/my-inference-component``.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide*.\n To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide*. Look for the ARN in the table at the bottom of the page.",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide*.\n To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide*. Look for the ARN in the table at the bottom of the page.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// RoleARN is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ScalableDimension
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.\n  +   ``ecs:service:DesiredCount`` - The desired task count of an ECS service.\n  +   ``elasticmapreduce:instancegroup:InstanceCount`` - The instance count of an EMR Instance Group.\n  +   ``ec2:spot-fleet-request:TargetCapacity`` - The target capacity of a Spot Fleet.\n  +   ``appstream:fleet:DesiredCapacity`` - The desired capacity of an AppStream 2.0 fleet.\n  +   ``dynamodb:table:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB table.\n  +   ``dynamodb:table:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB table.\n  +   ``dynamodb:index:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB global secondary index.\n  +   ``dynamodb:index:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB global secondary index.\n  +   ``rds:cluster:ReadReplicaCount`` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.\n  +   ``sagemaker:variant:DesiredInstanceCount`` - The number of EC2 instances for a SageMaker model endpoint variant.\n  +   ``custom-resource:ResourceType:Property`` - The scalable dimension for a custom resource provided by your own application or service.\n  +   ``comprehend:document-classifier-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend document classification endpoint.\n  +   ``comprehend:entity-recognizer-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.\n  +   ``lambda:function:ProvisionedConcurrency`` - The provisioned concurrency for a Lambda function.\n  +   ``cassandra:table:ReadCapacityUnits`` - The provisioned read capacity for an Amazon Keyspaces table.\n  +   ``cassandra:table:WriteCapacityUnits`` - The provisioned write capacity for an Amazon Keyspaces table.\n  +   ``kafka:broker-storage:VolumeSize`` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.\n  +   ``elasticache:replication-group:NodeGroups`` - The number of node groups for an Amazon ElastiCache replication group.\n  +   ``elasticache:replication-group:Replicas`` - The number of replicas per node group for an Amazon ElastiCache replication group.\n  +   ``neptune:cluster:ReadReplicaCount`` - The count of read replicas in an Amazon Neptune DB cluster.\n  +   ``sagemaker:variant:DesiredProvisionedConcurrency`` - The provisioned concurrency for a SageMaker serverless endpoint.\n  +   ``sagemaker:inference-component:DesiredCopyCount`` - The number of copies across an endpoint for a SageMaker inference component.",
		//	  "type": "string"
		//	}
		"scalable_dimension": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property.\n  +   ``ecs:service:DesiredCount`` - The desired task count of an ECS service.\n  +   ``elasticmapreduce:instancegroup:InstanceCount`` - The instance count of an EMR Instance Group.\n  +   ``ec2:spot-fleet-request:TargetCapacity`` - The target capacity of a Spot Fleet.\n  +   ``appstream:fleet:DesiredCapacity`` - The desired capacity of an AppStream 2.0 fleet.\n  +   ``dynamodb:table:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB table.\n  +   ``dynamodb:table:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB table.\n  +   ``dynamodb:index:ReadCapacityUnits`` - The provisioned read capacity for a DynamoDB global secondary index.\n  +   ``dynamodb:index:WriteCapacityUnits`` - The provisioned write capacity for a DynamoDB global secondary index.\n  +   ``rds:cluster:ReadReplicaCount`` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.\n  +   ``sagemaker:variant:DesiredInstanceCount`` - The number of EC2 instances for a SageMaker model endpoint variant.\n  +   ``custom-resource:ResourceType:Property`` - The scalable dimension for a custom resource provided by your own application or service.\n  +   ``comprehend:document-classifier-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend document classification endpoint.\n  +   ``comprehend:entity-recognizer-endpoint:DesiredInferenceUnits`` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.\n  +   ``lambda:function:ProvisionedConcurrency`` - The provisioned concurrency for a Lambda function.\n  +   ``cassandra:table:ReadCapacityUnits`` - The provisioned read capacity for an Amazon Keyspaces table.\n  +   ``cassandra:table:WriteCapacityUnits`` - The provisioned write capacity for an Amazon Keyspaces table.\n  +   ``kafka:broker-storage:VolumeSize`` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.\n  +   ``elasticache:replication-group:NodeGroups`` - The number of node groups for an Amazon ElastiCache replication group.\n  +   ``elasticache:replication-group:Replicas`` - The number of replicas per node group for an Amazon ElastiCache replication group.\n  +   ``neptune:cluster:ReadReplicaCount`` - The count of read replicas in an Amazon Neptune DB cluster.\n  +   ``sagemaker:variant:DesiredProvisionedConcurrency`` - The provisioned concurrency for a SageMaker serverless endpoint.\n  +   ``sagemaker:inference-component:DesiredCopyCount`` - The number of copies across an endpoint for a SageMaker inference component.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScheduledActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scheduled actions for the scalable target. Duplicates aren't allowed.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "``ScheduledAction`` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies a scheduled action for a scalable target. \n For more information, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide*.",
		//	    "properties": {
		//	      "EndTime": {
		//	        "description": "The date and time that the action is scheduled to end, in UTC.",
		//	        "type": "string"
		//	      },
		//	      "ScalableTargetAction": {
		//	        "additionalProperties": false,
		//	        "description": "The new minimum and maximum capacity. You can set both values or just one. At the scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling scales in to the maximum capacity.",
		//	        "properties": {
		//	          "MaxCapacity": {
		//	            "description": "The maximum capacity.",
		//	            "type": "integer"
		//	          },
		//	          "MinCapacity": {
		//	            "description": "The minimum capacity.",
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Schedule": {
		//	        "description": "The schedule for this action. The following formats are supported:\n  +  At expressions - \"``at(yyyy-mm-ddThh:mm:ss)``\"\n  +  Rate expressions - \"``rate(value unit)``\"\n  +  Cron expressions - \"``cron(fields)``\"\n  \n At expressions are useful for one-time schedules. Cron expressions are useful for scheduled actions that run periodically at a specified date and time, and rate expressions are useful for scheduled actions that run at a regular interval.\n At and cron expressions use Universal Coordinated Time (UTC) by default.\n The cron format consists of six fields separated by white spaces: [Minutes] [Hours] [Day_of_Month] [Month] [Day_of_Week] [Year].\n For rate expressions, *value* is a positive integer and *unit* is ``minute`` | ``minutes`` | ``hour`` | ``hours`` | ``day`` | ``days``.",
		//	        "type": "string"
		//	      },
		//	      "ScheduledActionName": {
		//	        "description": "The name of the scheduled action. This name must be unique among all other scheduled actions on the specified scalable target.",
		//	        "type": "string"
		//	      },
		//	      "StartTime": {
		//	        "description": "The date and time that the action is scheduled to begin, in UTC.",
		//	        "type": "string"
		//	      },
		//	      "Timezone": {
		//	        "description": "The time zone used when referring to the date and time of a scheduled action, when the scheduled action uses an at or cron expression.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ScheduledActionName",
		//	      "Schedule"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"scheduled_actions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EndTime
					"end_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The date and time that the action is scheduled to end, in UTC.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ScalableTargetAction
					"scalable_target_action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MaxCapacity
							"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The maximum capacity.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MinCapacity
							"min_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The minimum capacity.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The new minimum and maximum capacity. You can set both values or just one. At the scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling scales in to the maximum capacity.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Schedule
					"schedule": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The schedule for this action. The following formats are supported:\n  +  At expressions - \"``at(yyyy-mm-ddThh:mm:ss)``\"\n  +  Rate expressions - \"``rate(value unit)``\"\n  +  Cron expressions - \"``cron(fields)``\"\n  \n At expressions are useful for one-time schedules. Cron expressions are useful for scheduled actions that run periodically at a specified date and time, and rate expressions are useful for scheduled actions that run at a regular interval.\n At and cron expressions use Universal Coordinated Time (UTC) by default.\n The cron format consists of six fields separated by white spaces: [Minutes] [Hours] [Day_of_Month] [Month] [Day_of_Week] [Year].\n For rate expressions, *value* is a positive integer and *unit* is ``minute`` | ``minutes`` | ``hour`` | ``hours`` | ``day`` | ``days``.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ScheduledActionName
					"scheduled_action_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the scheduled action. This name must be unique among all other scheduled actions on the specified scalable target.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: StartTime
					"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The date and time that the action is scheduled to begin, in UTC.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Timezone
					"timezone": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The time zone used when referring to the date and time of a scheduled action, when the scheduled action uses an at or cron expression.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The scheduled actions for the scalable target. Duplicates aren't allowed.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceNamespace
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The namespace of the AWS service that provides the resource, or a ``custom-resource``.",
		//	  "type": "string"
		//	}
		"service_namespace": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The namespace of the AWS service that provides the resource, or a ``custom-resource``.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SuspendedState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to ``true`` suspends the specified scaling activities. Setting it to ``false`` (default) resumes the specified scaling activities. \n  *Suspension Outcomes* \n  +  For ``DynamicScalingInSuspended``, while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.\n  +  For ``DynamicScalingOutSuspended``, while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.\n  +  For ``ScheduledScalingSuspended``, while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.",
		//	  "properties": {
		//	    "DynamicScalingInSuspended": {
		//	      "description": "Whether scale in by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is ``false``.",
		//	      "type": "boolean"
		//	    },
		//	    "DynamicScalingOutSuspended": {
		//	      "description": "Whether scale out by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is ``false``.",
		//	      "type": "boolean"
		//	    },
		//	    "ScheduledScalingSuspended": {
		//	      "description": "Whether scheduled scaling is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is ``false``.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"suspended_state": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DynamicScalingInSuspended
				"dynamic_scaling_in_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether scale in by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is ``false``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DynamicScalingOutSuspended
				"dynamic_scaling_out_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether scale out by a target tracking scaling policy or a step scaling policy is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is ``false``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ScheduledScalingSuspended
				"scheduled_scaling_suspended": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether scheduled scaling is suspended. Set the value to ``true`` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is ``false``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to ``true`` suspends the specified scaling activities. Setting it to ``false`` (default) resumes the specified scaling activities. \n  *Suspension Outcomes* \n  +  For ``DynamicScalingInSuspended``, while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.\n  +  For ``DynamicScalingOutSuspended``, while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.\n  +  For ``ScheduledScalingSuspended``, while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "The ``AWS::ApplicationAutoScaling::ScalableTarget`` resource specifies a resource that Application Auto Scaling can scale, such as an AWS::DynamoDB::Table or AWS::ECS::Service resource.\n For more information, see [Getting started](https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the *Application Auto Scaling User Guide*.\n  If the resource that you want Application Auto Scaling to scale is not yet created in your account, add a dependency on the resource when registering it as a scalable target using the [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) attribute.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApplicationAutoScaling::ScalableTarget").WithTerraformTypeName("awscc_applicationautoscaling_scalable_target")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dynamic_scaling_in_suspended":  "DynamicScalingInSuspended",
		"dynamic_scaling_out_suspended": "DynamicScalingOutSuspended",
		"end_time":                      "EndTime",
		"max_capacity":                  "MaxCapacity",
		"min_capacity":                  "MinCapacity",
		"resource_id":                   "ResourceId",
		"role_arn":                      "RoleARN",
		"scalable_dimension":            "ScalableDimension",
		"scalable_target_action":        "ScalableTargetAction",
		"scalable_target_id":            "Id",
		"schedule":                      "Schedule",
		"scheduled_action_name":         "ScheduledActionName",
		"scheduled_actions":             "ScheduledActions",
		"scheduled_scaling_suspended":   "ScheduledScalingSuspended",
		"service_namespace":             "ServiceNamespace",
		"start_time":                    "StartTime",
		"suspended_state":               "SuspendedState",
		"timezone":                      "Timezone",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/RoleARN",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
