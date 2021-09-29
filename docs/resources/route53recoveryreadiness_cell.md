---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53recoveryreadiness_cell Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The API Schema for AWS Route53 Recovery Readiness Cells.
---

# awscc_route53recoveryreadiness_cell (Resource)

The API Schema for AWS Route53 Recovery Readiness Cells.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cell_name** (String) The name of the cell to create.

### Optional

- **cells** (List of String) A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.
- **tags** (Attributes List) A collection of tags associated with a resource (see [below for nested schema](#nestedatt--tags))

### Read-Only

- **cell_arn** (String) The Amazon Resource Name (ARN) of the cell.
- **id** (String) Uniquely identifies the resource.
- **parent_readiness_scopes** (List of String) The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- **key** (String)
- **value** (List of String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53recoveryreadiness_cell.example <resource ID>
```