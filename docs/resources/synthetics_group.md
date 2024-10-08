---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_synthetics_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Synthetics::Group
---

# awscc_synthetics_group (Resource)

Resource Type definition for AWS::Synthetics::Group



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the group.

### Optional

- `resource_arns` (List of String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `group_id` (String) Id of the group.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_synthetics_group.example "name"
```
