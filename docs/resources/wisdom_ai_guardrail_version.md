---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_wisdom_ai_guardrail_version Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Wisdom::AIGuardrailVersion Resource Type
---

# awscc_wisdom_ai_guardrail_version (Resource)

Definition of AWS::Wisdom::AIGuardrailVersion Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ai_guardrail_id` (String)
- `assistant_id` (String)

### Optional

- `modified_time_seconds` (Number)

### Read-Only

- `ai_guardrail_arn` (String)
- `ai_guardrail_version_id` (String)
- `assistant_arn` (String)
- `id` (String) Uniquely identifies the resource.
- `version_number` (Number)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_wisdom_ai_guardrail_version.example "assistant_id|ai_guardrail_id|version_number"
```