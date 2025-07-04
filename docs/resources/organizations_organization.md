---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_organizations_organization Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Organizations::Organization
---

# awscc_organizations_organization (Resource)

Resource schema for AWS::Organizations::Organization



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `feature_set` (String) Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of an organization.
- `id` (String) Uniquely identifies the resource.
- `management_account_arn` (String) The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization.
- `management_account_email` (String) The email address that is associated with the AWS account that is designated as the management account for the organization.
- `management_account_id` (String) The unique identifier (ID) of the management account of an organization.
- `organization_id` (String) The unique identifier (ID) of an organization.
- `root_id` (String) The unique identifier (ID) for the root.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_organizations_organization.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_organizations_organization.example "id"
```
