---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_kms_alias Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::KMS::Alias resource specifies a display name for a customer master key (CMK) in AWS Key Management Service (AWS KMS). You can use an alias to identify a CMK in cryptographic operations.
---

# awscc_kms_alias (Resource)

The AWS::KMS::Alias resource specifies a display name for a customer master key (CMK) in AWS Key Management Service (AWS KMS). You can use an alias to identify a CMK in cryptographic operations.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **alias_name** (String) Specifies the alias name. This value must begin with alias/ followed by a name, such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The alias/aws/ prefix is reserved for AWS managed CMKs.
- **target_key_id** (String) Identifies the CMK to which the alias refers. Specify the key ID or the Amazon Resource Name (ARN) of the CMK. You cannot specify another alias. For help finding the key ID and ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.

### Read-Only

- **id** (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_kms_alias.example <resource ID>
```