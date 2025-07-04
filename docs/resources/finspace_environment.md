
---
page_title: "awscc_finspace_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_finspace_environment (Resource)

An example resource schema demonstrating some basic constructs and validation rules.

## Example Usage

### Create FinSpace Environment with KMS Encryption

Creates a FinSpace environment with a dedicated KMS key for encryption and sets up an admin superuser, demonstrating secure environment creation with custom key policy for FinSpace service integration.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create KMS key for FinSpace environment
resource "awscc_kms_key" "finspace" {
  description = "KMS key for FinSpace environment"
  key_usage   = "ENCRYPT_DECRYPT"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow FinSpace Service"
        Effect = "Allow"
        Principal = {
          Service = "finspace.amazonaws.com"
        }
        Action = [
          "kms:Decrypt",
          "kms:GenerateDataKey",
          "kms:CreateGrant",
          "kms:RetireGrant",
          "kms:DescribeKey"
        ]
        Resource = "*"
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Example of FinSpace Environment
resource "awscc_finspace_environment" "example" {
  name        = "example-finspace-env"
  description = "Example FinSpace Environment created with AWSCC provider"
  kms_key_id  = awscc_kms_key.finspace.id

  # Example of superuser parameters
  superuser_parameters = {
    emailAddress = "admin@example.com"
    firstName    = "Admin"
    lastName     = "User"
  }

  # Example tags
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the Environment

### Optional

- `data_bundles` (List of String) ARNs of FinSpace Data Bundles to install
- `description` (String) Description of the Environment
- `federation_mode` (String) Federation mode used with the Environment
- `federation_parameters` (Attributes) Additional parameters to identify Federation mode (see [below for nested schema](#nestedatt--federation_parameters))
- `kms_key_id` (String) KMS key used to encrypt customer data within FinSpace Environment infrastructure
- `superuser_parameters` (Attributes) Parameters of the first Superuser for the FinSpace Environment (see [below for nested schema](#nestedatt--superuser_parameters))
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `aws_account_id` (String) AWS account ID associated with the Environment
- `dedicated_service_account_id` (String) ID for FinSpace created account used to store Environment artifacts
- `environment_arn` (String) ARN of the Environment
- `environment_id` (String) Unique identifier for representing FinSpace Environment
- `environment_url` (String) URL used to login to the Environment
- `id` (String) Uniquely identifies the resource.
- `sage_maker_studio_domain_url` (String) SageMaker Studio Domain URL associated with the Environment
- `status` (String) State of the Environment

<a id="nestedatt--federation_parameters"></a>
### Nested Schema for `federation_parameters`

Optional:

- `application_call_back_url` (String) SAML metadata URL to link with the Environment
- `attribute_map` (Attributes List) Attribute map for SAML configuration (see [below for nested schema](#nestedatt--federation_parameters--attribute_map))
- `federation_provider_name` (String) Federation provider name to link with the Environment
- `federation_urn` (String) SAML metadata URL to link with the Environment
- `saml_metadata_document` (String) SAML metadata document to link the federation provider to the Environment
- `saml_metadata_url` (String) SAML metadata URL to link with the Environment

<a id="nestedatt--federation_parameters--attribute_map"></a>
### Nested Schema for `federation_parameters.attribute_map`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.



<a id="nestedatt--superuser_parameters"></a>
### Nested Schema for `superuser_parameters`

Optional:

- `email_address` (String) Email address
- `first_name` (String) First name
- `last_name` (String) Last name


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_finspace_environment.example
  id = "environment_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_finspace_environment.example "environment_id"
```
