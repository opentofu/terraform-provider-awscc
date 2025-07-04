---
page_title: "awscc_timestream_database Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::Timestream::Database resource creates a Timestream database.
---

# awscc_timestream_database (Resource)

The AWS::Timestream::Database resource creates a Timestream database.

## Example Usage

### Create Database
Create a Timestream database
```terraform
resource "aws_kms_key" "this" {
  description = "Timestream KMS Key"
}

resource "awscc_timestream_database" "this" {
  database_name = "MyTimestreamDB"
  kms_key_id    = aws_kms_key.this.key_id
  tags = [{
    key   = "Managed By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `database_name` (String) The name for the database. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the database name.
- `kms_key_id` (String) The KMS key for the database. If the KMS key is not specified, the database will be encrypted with a Timestream managed KMS key located in your account.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_timestream_database.example
  id = "database_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_timestream_database.example "database_name"
```
