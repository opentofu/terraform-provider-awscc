---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_quicksight_analysis Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of the AWS::QuickSight::Analysis Resource Type.
---

# awscc_quicksight_analysis (Resource)

Definition of the AWS::QuickSight::Analysis Resource Type.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `analysis_id` (String)
- `aws_account_id` (String)
- `source_entity` (Attributes) <p>The source entity of an analysis.</p> (see [below for nested schema](#nestedatt--source_entity))

### Optional

- `errors` (Attributes List) <p>Errors associated with the analysis.</p> (see [below for nested schema](#nestedatt--errors))
- `name` (String) <p>The descriptive name of the analysis.</p>
- `parameters` (Attributes) <p>A list of QuickSight parameters and the list's override values.</p> (see [below for nested schema](#nestedatt--parameters))
- `permissions` (Attributes List) <p>A structure that describes the principals and the resource-level permissions on an
            analysis. You can use the <code>Permissions</code> structure to grant permissions by
            providing a list of AWS Identity and Access Management (IAM) action information for each
            principal listed by Amazon Resource Name (ARN). </p>

        <p>To specify no permissions, omit <code>Permissions</code>.</p> (see [below for nested schema](#nestedatt--permissions))
- `tags` (Attributes List) <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the
            analysis.</p> (see [below for nested schema](#nestedatt--tags))
- `theme_arn` (String) <p>The ARN of the theme of the analysis.</p>

### Read-Only

- `arn` (String) <p>The Amazon Resource Name (ARN) of the analysis.</p>
- `created_time` (String) <p>The time that the analysis was created.</p>
- `data_set_arns` (List of String) <p>The ARNs of the datasets of the analysis.</p>
- `id` (String) Uniquely identifies the resource.
- `last_updated_time` (String) <p>The time that the analysis was last updated.</p>
- `sheets` (Attributes List) <p>A list of the associated sheets with the unique identifier and name of each sheet.</p> (see [below for nested schema](#nestedatt--sheets))
- `status` (String)

<a id="nestedatt--source_entity"></a>
### Nested Schema for `source_entity`

Optional:

- `source_template` (Attributes) <p>The source template of an analysis.</p> (see [below for nested schema](#nestedatt--source_entity--source_template))

<a id="nestedatt--source_entity--source_template"></a>
### Nested Schema for `source_entity.source_template`

Optional:

- `arn` (String) <p>The Amazon Resource Name (ARN) of the source template of an analysis.</p>
- `data_set_references` (Attributes List) <p>The dataset references of the source template of an analysis.</p> (see [below for nested schema](#nestedatt--source_entity--source_template--data_set_references))

<a id="nestedatt--source_entity--source_template--data_set_references"></a>
### Nested Schema for `source_entity.source_template.data_set_references`

Optional:

- `data_set_arn` (String) <p>Dataset Amazon Resource Name (ARN).</p>
- `data_set_placeholder` (String) <p>Dataset placeholder.</p>




<a id="nestedatt--errors"></a>
### Nested Schema for `errors`

Optional:

- `message` (String) <p>The message associated with the analysis error.</p>
- `type` (String)


<a id="nestedatt--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `date_time_parameters` (Attributes List) <p>Date-time parameters.</p> (see [below for nested schema](#nestedatt--parameters--date_time_parameters))
- `decimal_parameters` (Attributes List) <p>Decimal parameters.</p> (see [below for nested schema](#nestedatt--parameters--decimal_parameters))
- `integer_parameters` (Attributes List) <p>Integer parameters.</p> (see [below for nested schema](#nestedatt--parameters--integer_parameters))
- `string_parameters` (Attributes List) <p>String parameters.</p> (see [below for nested schema](#nestedatt--parameters--string_parameters))

<a id="nestedatt--parameters--date_time_parameters"></a>
### Nested Schema for `parameters.date_time_parameters`

Optional:

- `name` (String) <p>A display name for the date-time parameter.</p>
- `values` (List of String) <p>The values for the date-time parameter.</p>


<a id="nestedatt--parameters--decimal_parameters"></a>
### Nested Schema for `parameters.decimal_parameters`

Optional:

- `name` (String) <p>A display name for the decimal parameter.</p>
- `values` (List of Number) <p>The values for the decimal parameter.</p>


<a id="nestedatt--parameters--integer_parameters"></a>
### Nested Schema for `parameters.integer_parameters`

Optional:

- `name` (String) <p>The name of the integer parameter.</p>
- `values` (List of Number) <p>The values for the integer parameter.</p>


<a id="nestedatt--parameters--string_parameters"></a>
### Nested Schema for `parameters.string_parameters`

Optional:

- `name` (String) <p>A display name for a string parameter.</p>
- `values` (List of String) <p>The values of a string parameter.</p>



<a id="nestedatt--permissions"></a>
### Nested Schema for `permissions`

Optional:

- `actions` (List of String) <p>The IAM action to grant or revoke permissions on.</p>
- `principal` (String) <p>The Amazon Resource Name (ARN) of the principal. This can be one of the
            following:</p>
        <ul>
            <li>
                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>
            </li>
            <li>
                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>
            </li>
            <li>
                <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight
                    ARN. Use this option only to share resources (templates) across AWS accounts.
                    (This is less common.) </p>
            </li>
         </ul>


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) <p>Tag key.</p>
- `value` (String) <p>Tag value.</p>


<a id="nestedatt--sheets"></a>
### Nested Schema for `sheets`

Read-Only:

- `name` (String) <p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight
            console.</p>
- `sheet_id` (String) <p>The unique identifier associated with a sheet.</p>

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_quicksight_analysis.example
  id = "analysis_id|aws_account_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_quicksight_analysis.example "analysis_id|aws_account_id"
```
