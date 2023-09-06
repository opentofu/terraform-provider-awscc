---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_config_config_rule Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for AWS Config ConfigRule
---

# awscc_config_config_rule (Resource)

Schema for AWS Config ConfigRule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source` (Attributes) Source of events for the AWS Config rule (see [below for nested schema](#nestedatt--source))

### Optional

- `compliance` (Attributes) Compliance details of the Config rule (see [below for nested schema](#nestedatt--compliance))
- `config_rule_name` (String) Name for the AWS Config rule
- `description` (String) Description provided for the AWS Config rule
- `evaluation_modes` (Attributes List) List of EvaluationModeConfiguration objects (see [below for nested schema](#nestedatt--evaluation_modes))
- `input_parameters` (String) JSON string passed the Lambda function
- `maximum_execution_frequency` (String) Maximum frequency at which the rule has to be evaluated
- `scope` (Attributes) Scope to constrain which resources can trigger the AWS Config rule (see [below for nested schema](#nestedatt--scope))

### Read-Only

- `arn` (String) ARN generated for the AWS Config rule
- `config_rule_id` (String) ID of the config rule
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--source"></a>
### Nested Schema for `source`

Required:

- `owner` (String) Owner of the config rule

Optional:

- `custom_policy_details` (Attributes) Custom policy details when rule is custom owned (see [below for nested schema](#nestedatt--source--custom_policy_details))
- `source_details` (Attributes List) List of message types that can trigger the rule (see [below for nested schema](#nestedatt--source--source_details))
- `source_identifier` (String) Identifier for the source of events

<a id="nestedatt--source--custom_policy_details"></a>
### Nested Schema for `source.custom_policy_details`

Optional:

- `enable_debug_log_delivery` (Boolean) Logging toggle for custom policy rule
- `policy_runtime` (String) Runtime system for custom policy rule
- `policy_text` (String) Policy definition containing logic for custom policy rule


<a id="nestedatt--source--source_details"></a>
### Nested Schema for `source.source_details`

Required:

- `event_source` (String) Source of event that can trigger the rule
- `message_type` (String) Notification type that can trigger the rule

Optional:

- `maximum_execution_frequency` (String) Frequency at which the rule has to be evaluated



<a id="nestedatt--compliance"></a>
### Nested Schema for `compliance`

Read-Only:

- `type` (String) Compliance type determined by the Config rule


<a id="nestedatt--evaluation_modes"></a>
### Nested Schema for `evaluation_modes`

Optional:

- `mode` (String) Mode of evaluation of AWS Config rule


<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Optional:

- `compliance_resource_id` (String) ID of the only one resource which we want to trigger the rule
- `compliance_resource_types` (List of String) Resource types of resources which we want to trigger the rule
- `tag_key` (String) Tag key applied only to resources which we want to trigger the rule
- `tag_value` (String) Tag value applied only to resources which we want to trigger the rule

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_config_config_rule.example <resource ID>
```