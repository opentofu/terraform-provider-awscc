---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_event_bridge_rule_template_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaLive::EventBridgeRuleTemplateGroup
---

# awscc_medialive_event_bridge_rule_template_group (Data Source)

Data Source schema for AWS::MediaLive::EventBridgeRuleTemplateGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) An eventbridge rule template group's ARN (Amazon Resource Name)
- `created_at` (String)
- `description` (String) A resource's optional description.
- `event_bridge_rule_template_group_id` (String) An eventbridge rule template group's id. AWS provided template groups have ids that start with `aws-`
- `identifier` (String)
- `modified_at` (String)
- `name` (String) A resource's name. Names must be unique within the scope of a resource type in a specific region.
- `tags` (Map of String) Represents the tags associated with a resource.
