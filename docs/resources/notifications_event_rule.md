---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_notifications_event_rule Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Notifications::EventRule
---

# awscc_notifications_event_rule (Resource)

Resource Type definition for AWS::Notifications::EventRule



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_type` (String)
- `notification_configuration_arn` (String)
- `regions` (List of String)
- `source` (String)

### Optional

- `event_pattern` (String)

### Read-Only

- `arn` (String)
- `creation_time` (String)
- `id` (String) Uniquely identifies the resource.
- `managed_rules` (List of String)
- `status_summary_by_region` (Attributes Map) (see [below for nested schema](#nestedatt--status_summary_by_region))

<a id="nestedatt--status_summary_by_region"></a>
### Nested Schema for `status_summary_by_region`

Read-Only:

- `reason` (String)
- `status` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_notifications_event_rule.example "arn"
```
