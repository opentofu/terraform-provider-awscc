---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_redshiftserverless_workgroup Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::RedshiftServerless::Workgroup
---

# awscc_redshiftserverless_workgroup (Data Source)

Data Source schema for AWS::RedshiftServerless::Workgroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `base_capacity` (Number) The base compute capacity of the workgroup in Redshift Processing Units (RPUs).
- `config_parameters` (Attributes Set) A list of parameters to set for finer control over a database. Available options are datestyle, enable_user_activity_logging, query_group, search_path, max_query_execution_time, and require_ssl. (see [below for nested schema](#nestedatt--config_parameters))
- `enhanced_vpc_routing` (Boolean) The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
- `max_capacity` (Number) The max compute capacity of the workgroup in Redshift Processing Units (RPUs).
- `namespace_name` (String) The namespace the workgroup is associated with.
- `port` (Number) The custom port to use when connecting to a workgroup. Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.
- `price_performance_target` (Attributes) A property that represents the price performance target settings for the workgroup. (see [below for nested schema](#nestedatt--price_performance_target))
- `publicly_accessible` (Boolean) A value that specifies whether the workgroup can be accessible from a public network.
- `recovery_point_id` (String) The recovery point id to restore from.
- `security_group_ids` (List of String) A list of security group IDs to associate with the workgroup.
- `snapshot_arn` (String) The Amazon Resource Name (ARN) of the snapshot to restore from.
- `snapshot_name` (String) The snapshot name to restore from.
- `snapshot_owner_account` (String) The Amazon Web Services account that owns the snapshot.
- `subnet_ids` (List of String) A list of subnet IDs the workgroup is associated with.
- `tags` (Attributes List) The map of the key-value pairs used to tag the workgroup. (see [below for nested schema](#nestedatt--tags))
- `track_name` (String)
- `workgroup` (Attributes) Definition for workgroup resource (see [below for nested schema](#nestedatt--workgroup))
- `workgroup_name` (String) The name of the workgroup.

<a id="nestedatt--config_parameters"></a>
### Nested Schema for `config_parameters`

Read-Only:

- `parameter_key` (String)
- `parameter_value` (String)


<a id="nestedatt--price_performance_target"></a>
### Nested Schema for `price_performance_target`

Read-Only:

- `level` (Number)
- `status` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


<a id="nestedatt--workgroup"></a>
### Nested Schema for `workgroup`

Read-Only:

- `base_capacity` (Number)
- `config_parameters` (Attributes Set) (see [below for nested schema](#nestedatt--workgroup--config_parameters))
- `creation_date` (String)
- `endpoint` (Attributes) (see [below for nested schema](#nestedatt--workgroup--endpoint))
- `enhanced_vpc_routing` (Boolean)
- `max_capacity` (Number)
- `namespace_name` (String)
- `price_performance_target` (Attributes) (see [below for nested schema](#nestedatt--workgroup--price_performance_target))
- `publicly_accessible` (Boolean)
- `security_group_ids` (List of String)
- `status` (String)
- `subnet_ids` (List of String)
- `track_name` (String)
- `workgroup_arn` (String)
- `workgroup_id` (String)
- `workgroup_name` (String)

<a id="nestedatt--workgroup--config_parameters"></a>
### Nested Schema for `workgroup.config_parameters`

Read-Only:

- `parameter_key` (String)
- `parameter_value` (String)


<a id="nestedatt--workgroup--endpoint"></a>
### Nested Schema for `workgroup.endpoint`

Read-Only:

- `address` (String)
- `port` (Number)
- `vpc_endpoints` (Attributes List) (see [below for nested schema](#nestedatt--workgroup--endpoint--vpc_endpoints))

<a id="nestedatt--workgroup--endpoint--vpc_endpoints"></a>
### Nested Schema for `workgroup.endpoint.vpc_endpoints`

Read-Only:

- `network_interfaces` (Attributes List) (see [below for nested schema](#nestedatt--workgroup--endpoint--vpc_endpoints--network_interfaces))
- `vpc_endpoint_id` (String)
- `vpc_id` (String)

<a id="nestedatt--workgroup--endpoint--vpc_endpoints--network_interfaces"></a>
### Nested Schema for `workgroup.endpoint.vpc_endpoints.network_interfaces`

Read-Only:

- `availability_zone` (String)
- `network_interface_id` (String)
- `private_ip_address` (String)
- `subnet_id` (String)




<a id="nestedatt--workgroup--price_performance_target"></a>
### Nested Schema for `workgroup.price_performance_target`

Read-Only:

- `level` (Number)
- `status` (String)
