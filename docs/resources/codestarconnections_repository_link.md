
---
page_title: "awscc_codestarconnections_repository_link Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.
---

# awscc_codestarconnections_repository_link (Resource)

Schema for AWS::CodeStarConnections::RepositoryLink resource which is used to aggregate repository metadata relevant to synchronizing source provider content to AWS Resources.

## Example Usage

### CodeStar Repository Link with GitHub Connection

Creates a CodeStar repository link that connects a GitHub repository to AWS using a CodeStar connection, enabling seamless integration between GitHub and AWS services.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create the CodeStar Connection first
resource "awscc_codestarconnections_connection" "example" {
  connection_name = "example-connection"
  provider_type   = "GitHub"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create the Repository Link
resource "awscc_codestarconnections_repository_link" "example" {
  connection_arn  = awscc_codestarconnections_connection.example.connection_arn
  owner_id        = "example-owner"
  repository_name = "example-repo"
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_arn` (String) The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services.
- `owner_id` (String) the ID of the entity that owns the repository.
- `repository_name` (String) The repository for which the link is being created.

### Optional

- `encryption_key_arn` (String) The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used.
- `tags` (Attributes List) Specifies the tags applied to a RepositoryLink. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `provider_type` (String) The name of the external provider where your third-party code repository is configured.
- `repository_link_arn` (String) A unique Amazon Resource Name (ARN) to designate the repository link.
- `repository_link_id` (String) A UUID that uniquely identifies the RepositoryLink.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_codestarconnections_repository_link.example
  id = "repository_link_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_codestarconnections_repository_link.example "repository_link_arn"
```
