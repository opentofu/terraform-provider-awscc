---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_identitystore_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::IdentityStore::Group
---

# awscc_identitystore_group (Resource)

Resource Type definition for AWS::IdentityStore::Group



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `display_name` (String) A string containing the name of the group. This value is commonly displayed when the group is referenced.
- `identity_store_id` (String) The globally unique identifier for the identity store.

### Optional

- `description` (String) A string containing the description of the group.

### Read-Only

- `group_id` (String) The unique identifier for a group in the identity store.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_identitystore_group.example "group_id|identity_store_id"
```
