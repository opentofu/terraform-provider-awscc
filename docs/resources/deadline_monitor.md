---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_deadline_monitor Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Deadline::Monitor Resource Type
---

# awscc_deadline_monitor (Resource)

Definition of AWS::Deadline::Monitor Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String)
- `identity_center_instance_arn` (String)
- `role_arn` (String)
- `subdomain` (String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `identity_center_application_arn` (String)
- `monitor_id` (String)
- `url` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_deadline_monitor.example "arn"
```
