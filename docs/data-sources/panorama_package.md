---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_panorama_package Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Panorama::Package
---

# awscc_panorama_package (Data Source)

Data Source schema for AWS::Panorama::Package



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String)
- `created_time` (Number)
- `package_id` (String)
- `package_name` (String) A name for the package.
- `storage_location` (Attributes) A storage location. (see [below for nested schema](#nestedatt--storage_location))
- `tags` (Attributes Set) Tags for the package. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--storage_location"></a>
### Nested Schema for `storage_location`

Read-Only:

- `binary_prefix_location` (String) The location's binary prefix.
- `bucket` (String) The location's bucket.
- `generated_prefix_location` (String) The location's generated prefix.
- `manifest_prefix_location` (String) The location's manifest prefix.
- `repo_prefix_location` (String) The location's repo prefix.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
