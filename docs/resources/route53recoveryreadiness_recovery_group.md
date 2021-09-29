---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoveryreadiness_recovery_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.
---

# awscc_route53recoveryreadiness_recovery_group (Resource)

AWS Route53 Recovery Readiness Recovery Group Schema and API specifications.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **recovery_group_name** (String) The name of the recovery group to create.

### Optional

- **cells** (List of String) A list of the cell Amazon Resource Names (ARNs) in the recovery group.
- **tags** (Attributes List) A collection of tags associated with a resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **id** (String) Uniquely identifies the resource.
- **recovery_group_arn** (String) A collection of tags associated with a resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53recoveryreadiness_recovery_group.example <resource ID>
```