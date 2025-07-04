---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_detective_member_invitation Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Detective::MemberInvitation
---

# awscc_detective_member_invitation (Resource)

Resource schema for AWS::Detective::MemberInvitation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `graph_arn` (String) The ARN of the graph to which the member account will be invited
- `member_email_address` (String) The root email address for the account to be invited, for validation. Updating this field has no effect.
- `member_id` (String) The AWS account ID to be invited to join the graph as a member

### Optional

- `disable_email_notification` (Boolean) When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.
- `message` (String) A message to be included in the email invitation sent to the invited account. Updating this field has no effect.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_detective_member_invitation.example
  id = "graph_arn|member_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_detective_member_invitation.example "graph_arn|member_id"
```
