
---
page_title: "awscc_notifications_managed_notification_additional_channel_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Notifications::ManagedNotificationAdditionalChannelAssociation
---

# awscc_notifications_managed_notification_additional_channel_association (Resource)

Resource Type definition for AWS::Notifications::ManagedNotificationAdditionalChannelAssociation

## Example Usage

### AWS Notification Channel Association with Slack

Associates an AWS Chatbot Slack channel configuration with AWS managed notifications to receive AWS Health billing alerts through Slack.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_caller_identity" "current" {}

# Create an AWS Chatbot Slack channel configuration first
resource "awscc_chatbot_slack_channel_configuration" "example" {
  configuration_name = "example-channel"
  slack_channel_id   = "EXAMPLE123" # Replace with actual Slack channel ID
  slack_workspace_id = "T0123456"   # Replace with actual Slack workspace ID

  iam_role_arn = awscc_iam_role.chatbot_role.arn

  guardrail_policies = [
    "arn:aws:iam::aws:policy/ReadOnlyAccess"
  ]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create IAM role for AWS Chatbot
resource "awscc_iam_role" "chatbot_role" {
  role_name = "AWSChatbotRole"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "chatbot.amazonaws.com"
        }
      }
    ]
  })
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Associate the channel with managed notification
resource "awscc_notifications_managed_notification_additional_channel_association" "example" {
  channel_arn                            = awscc_chatbot_slack_channel_configuration.example.arn
  managed_notification_configuration_arn = "arn:aws:notifications::${data.aws_caller_identity.current.account_id}:managed-notification-configuration/category/AWS-Health/sub-category/Billing"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `channel_arn` (String) ARN identifier of the channel.
Example: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops
- `managed_notification_configuration_arn` (String) ARN identifier of the Managed Notification.
Example: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_notifications_managed_notification_additional_channel_association.example
  id = "channel_arn|managed_notification_configuration_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_notifications_managed_notification_additional_channel_association.example "channel_arn|managed_notification_configuration_arn"
```
