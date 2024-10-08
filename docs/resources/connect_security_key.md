---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_security_key Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Connect::SecurityKey
---

# awscc_connect_security_key (Resource)

Resource Type definition for AWS::Connect::SecurityKey



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_id` (String) Amazon Connect instance identifier
- `key` (String) A valid security key in PEM format.

### Read-Only

- `association_id` (String) An associationID is automatically generated when a storage config is associated with an instance
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_connect_security_key.example "instance_id|association_id"
```
