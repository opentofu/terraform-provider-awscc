---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_comprehend_document_classifier Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Document Classifier enables training document classifier models.
---

# awscc_comprehend_document_classifier (Resource)

Document Classifier enables training document classifier models.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `data_access_role_arn` (String)
- `document_classifier_name` (String)
- `input_data_config` (Attributes) (see [below for nested schema](#nestedatt--input_data_config))
- `language_code` (String)

### Optional

- `mode` (String)
- `model_kms_key_id` (String)
- `model_policy` (String)
- `output_data_config` (Attributes) (see [below for nested schema](#nestedatt--output_data_config))
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))
- `version_name` (String)
- `volume_kms_key_id` (String)
- `vpc_config` (Attributes) (see [below for nested schema](#nestedatt--vpc_config))

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--input_data_config"></a>
### Nested Schema for `input_data_config`

Optional:

- `augmented_manifests` (Attributes Set) (see [below for nested schema](#nestedatt--input_data_config--augmented_manifests))
- `data_format` (String)
- `document_reader_config` (Attributes) (see [below for nested schema](#nestedatt--input_data_config--document_reader_config))
- `document_type` (String)
- `documents` (Attributes) (see [below for nested schema](#nestedatt--input_data_config--documents))
- `label_delimiter` (String)
- `s3_uri` (String)
- `test_s3_uri` (String)

<a id="nestedatt--input_data_config--augmented_manifests"></a>
### Nested Schema for `input_data_config.augmented_manifests`

Optional:

- `attribute_names` (Set of String)
- `s3_uri` (String)
- `split` (String)


<a id="nestedatt--input_data_config--document_reader_config"></a>
### Nested Schema for `input_data_config.document_reader_config`

Optional:

- `document_read_action` (String)
- `document_read_mode` (String)
- `feature_types` (Set of String)


<a id="nestedatt--input_data_config--documents"></a>
### Nested Schema for `input_data_config.documents`

Optional:

- `s3_uri` (String)
- `test_s3_uri` (String)



<a id="nestedatt--output_data_config"></a>
### Nested Schema for `output_data_config`

Optional:

- `kms_key_id` (String)
- `s3_uri` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--vpc_config"></a>
### Nested Schema for `vpc_config`

Optional:

- `security_group_ids` (Set of String)
- `subnets` (Set of String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_comprehend_document_classifier.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_comprehend_document_classifier.example "arn"
```
