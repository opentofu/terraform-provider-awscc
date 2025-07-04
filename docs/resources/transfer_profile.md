
---
page_title: "awscc_transfer_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Transfer::Profile
---

# awscc_transfer_profile (Resource)

Resource Type definition for AWS::Transfer::Profile

## Example Usage

### Create Transfer Profile for AS2 Protocol

Creates a local Transfer Profile for AS2 protocol with profile type set to LOCAL and specifies an AS2 identifier, with optional support for certificate management.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create the Transfer Profile
resource "awscc_transfer_profile" "example" {
  as_2_id      = "ExampleAs2Id"
  profile_type = "LOCAL"

  # Optional: Add certificates if needed
  # certificate_ids = ["certificate-id-1", "certificate-id-2"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_2_id` (String) AS2 identifier agreed with a trading partner.
- `profile_type` (String) Enum specifying whether the profile is local or associated with a trading partner.

### Optional

- `certificate_ids` (List of String) List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) Specifies the unique Amazon Resource Name (ARN) for the profile.
- `id` (String) Uniquely identifies the resource.
- `profile_id` (String) A unique identifier for the profile

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The name assigned to the tag that you create.
- `value` (String) Contains one or more values that you assigned to the key name you create.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_transfer_profile.example
  id = "profile_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_transfer_profile.example "profile_id"
```
