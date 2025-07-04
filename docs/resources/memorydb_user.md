
---
page_title: "awscc_memorydb_user Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::MemoryDB::User
---

# awscc_memorydb_user (Resource)

Resource Type definition for AWS::MemoryDB::User

## Example Usage

### Create MemoryDB User with Password Authentication

Creates a MemoryDB user with password-based authentication and full access permissions to all commands and keys in the database.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
resource "awscc_memorydb_user" "example" {
  user_name     = "sample-user"
  access_string = "on ~* +@all" # Full access to all commands and keys
  authentication_mode = {
    type      = "password"
    passwords = ["MyP@ssw0rd123!456789"] # You should use sensitive values in production
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `user_name` (String) The name of the user.

### Optional

- `access_string` (String) Access permissions string used for this user account.
- `authentication_mode` (Attributes) (see [below for nested schema](#nestedatt--authentication_mode))
- `tags` (Attributes Set) An array of key-value pairs to apply to this user. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the user account.
- `id` (String) Uniquely identifies the resource.
- `status` (String) Indicates the user status. Can be "active", "modifying" or "deleting".

<a id="nestedatt--authentication_mode"></a>
### Nested Schema for `authentication_mode`

Optional:

- `passwords` (List of String) Passwords used for this user account. You can create up to two passwords for each user.
- `type` (String) Type of authentication strategy for this user.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with 'aws:'. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_memorydb_user.example
  id = "user_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_memorydb_user.example "user_name"
```
