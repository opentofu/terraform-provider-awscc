---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_dsql_cluster Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::DSQL::Cluster
---

# awscc_dsql_cluster (Data Source)

Data Source schema for AWS::DSQL::Cluster



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `creation_time` (String) The time of when the cluster was created in ISO-8601 format.
- `deletion_protection_enabled` (Boolean) Whether deletion protection is enabled in this cluster.
- `encryption_details` (Attributes) The encryption configuration details for the cluster. (see [below for nested schema](#nestedatt--encryption_details))
- `identifier` (String) The ID of the created cluster.
- `kms_encryption_key` (String) The KMS key that encrypts data on the cluster.
- `multi_region_properties` (Attributes) The Multi-region properties associated to this cluster. (see [below for nested schema](#nestedatt--multi_region_properties))
- `resource_arn` (String) The Amazon Resource Name (ARN) for the cluster.
- `status` (String) The status of the cluster.
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `vpc_endpoint_service_name` (String) The VPC endpoint service name.

<a id="nestedatt--encryption_details"></a>
### Nested Schema for `encryption_details`

Read-Only:

- `encryption_status` (String) The status of encryption for the cluster.
- `encryption_type` (String) The type of encryption that protects data in the cluster.
- `kms_key_arn` (String) The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.


<a id="nestedatt--multi_region_properties"></a>
### Nested Schema for `multi_region_properties`

Read-Only:

- `clusters` (Set of String)
- `witness_region` (String) The witness region in a multi-region cluster.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
