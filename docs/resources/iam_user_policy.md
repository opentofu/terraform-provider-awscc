---
page_title: "awscc_iam_user_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Adds or updates an inline policy document that is embedded in the specified IAM user.
  An IAM user can also have a managed policy attached to it. To attach a managed policy to a user, use AWS::IAM::User https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html. To create a new managed policy, use AWS::IAM::ManagedPolicy https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html. For information about policies, see Managed policies and inline policies https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html in the IAM User Guide.
  For information about the maximum number of inline policies that you can embed in a user, see IAM and quotas https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html in the IAM User Guide.
---

# awscc_iam_user_policy (Resource)

Adds or updates an inline policy document that is embedded in the specified IAM user.
 An IAM user can also have a managed policy attached to it. To attach a managed policy to a user, use [AWS::IAM::User](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html). To create a new managed policy, use [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html). For information about policies, see [Managed policies and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide*.
 For information about the maximum number of inline policies that you can embed in a user, see [IAM and quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the *IAM User Guide*.

## Example Usage

### IAM user embedded inline policy document

The following example creates inline policy and attaches it to an IAM user

```terraform
resource "awscc_iam_user_policy" "example" {
  policy_name = "sample_iam_user_policy"
  user_name   = awscc_iam_user.example.id

  policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:ListAllMyBuckets",
          "s3:GetBucketLocation",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:s3:::*"
      },
    ]
  })
}

resource "awscc_iam_user" "example" {
  user_name = "sample_iam_user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy_name` (String) The name of the policy document.
 This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
- `user_name` (String) The name of the user to associate the policy with.
 This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-

### Optional

- `policy_document` (String) The policy document.
 You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.
 The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:
  +  Any printable ASCII character ranging from the space character (``\u0020``) through the end of the ASCII character range
  +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\u00FF``)
  +  The special characters tab (``\u0009``), line feed (``\u000A``), and carriage return (``\u000D``)

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_iam_user_policy.example
  id = "policy_name|user_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_iam_user_policy.example "policy_name|user_name"
```