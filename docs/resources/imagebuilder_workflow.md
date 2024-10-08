---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_imagebuilder_workflow Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::ImageBuilder::Workflow
---

# awscc_imagebuilder_workflow (Resource)

Resource schema for AWS::ImageBuilder::Workflow



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the workflow.
- `type` (String) The type of the workflow denotes whether the workflow is used to build, test, or distribute.
- `version` (String) The version of the workflow.

### Optional

- `change_description` (String) The change description of the workflow.
- `data` (String) The data of the workflow.
- `description` (String) The description of the workflow.
- `kms_key_id` (String) The KMS key identifier used to encrypt the workflow.
- `tags` (Map of String) The tags associated with the workflow.
- `uri` (String) The uri of the workflow.

### Read-Only

- `arn` (String) The Amazon Resource Name (ARN) of the workflow.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_imagebuilder_workflow.example "arn"
```
