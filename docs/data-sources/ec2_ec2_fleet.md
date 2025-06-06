---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_ec2_fleet Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::EC2Fleet
---

# awscc_ec2_ec2_fleet (Data Source)

Data Source schema for AWS::EC2::EC2Fleet



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `context` (String)
- `excess_capacity_termination_policy` (String)
- `fleet_id` (String)
- `launch_template_configs` (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs))
- `on_demand_options` (Attributes) (see [below for nested schema](#nestedatt--on_demand_options))
- `replace_unhealthy_instances` (Boolean)
- `spot_options` (Attributes) (see [below for nested schema](#nestedatt--spot_options))
- `tag_specifications` (Attributes List) (see [below for nested schema](#nestedatt--tag_specifications))
- `target_capacity_specification` (Attributes) (see [below for nested schema](#nestedatt--target_capacity_specification))
- `terminate_instances_with_expiration` (Boolean)
- `type` (String)
- `valid_from` (String)
- `valid_until` (String)

<a id="nestedatt--launch_template_configs"></a>
### Nested Schema for `launch_template_configs`

Read-Only:

- `launch_template_specification` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--launch_template_specification))
- `overrides` (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs--overrides))

<a id="nestedatt--launch_template_configs--launch_template_specification"></a>
### Nested Schema for `launch_template_configs.launch_template_specification`

Read-Only:

- `launch_template_id` (String)
- `launch_template_name` (String)
- `version` (String)


<a id="nestedatt--launch_template_configs--overrides"></a>
### Nested Schema for `launch_template_configs.overrides`

Read-Only:

- `availability_zone` (String)
- `block_device_mappings` (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--block_device_mappings))
- `instance_requirements` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements))
- `instance_type` (String)
- `max_price` (String)
- `placement` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--placement))
- `priority` (Number)
- `subnet_id` (String)
- `weighted_capacity` (Number)

<a id="nestedatt--launch_template_configs--overrides--block_device_mappings"></a>
### Nested Schema for `launch_template_configs.overrides.block_device_mappings`

Read-Only:

- `device_name` (String)
- `ebs` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--block_device_mappings--ebs))
- `no_device` (String)
- `virtual_name` (String)

<a id="nestedatt--launch_template_configs--overrides--block_device_mappings--ebs"></a>
### Nested Schema for `launch_template_configs.overrides.block_device_mappings.ebs`

Read-Only:

- `delete_on_termination` (Boolean)
- `encrypted` (Boolean)
- `iops` (Number)
- `kms_key_id` (String)
- `snapshot_id` (String)
- `volume_size` (Number)
- `volume_type` (String)



<a id="nestedatt--launch_template_configs--overrides--instance_requirements"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements`

Read-Only:

- `accelerator_count` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--accelerator_count))
- `accelerator_manufacturers` (List of String)
- `accelerator_names` (List of String)
- `accelerator_total_memory_mi_b` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--accelerator_total_memory_mi_b))
- `accelerator_types` (List of String)
- `allowed_instance_types` (List of String)
- `bare_metal` (String)
- `baseline_ebs_bandwidth_mbps` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--baseline_ebs_bandwidth_mbps))
- `baseline_performance_factors` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors))
- `burstable_performance` (String)
- `cpu_manufacturers` (List of String)
- `excluded_instance_types` (List of String)
- `instance_generations` (List of String)
- `local_storage` (String)
- `local_storage_types` (List of String)
- `max_spot_price_as_percentage_of_optimal_on_demand_price` (Number)
- `memory_gi_b_per_v_cpu` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--memory_gi_b_per_v_cpu))
- `memory_mi_b` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--memory_mi_b))
- `network_bandwidth_gbps` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--network_bandwidth_gbps))
- `network_interface_count` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--network_interface_count))
- `on_demand_max_price_percentage_over_lowest_price` (Number)
- `require_hibernate_support` (Boolean)
- `spot_max_price_percentage_over_lowest_price` (Number)
- `total_local_storage_gb` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--total_local_storage_gb))
- `v_cpu_count` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--v_cpu_count))

<a id="nestedatt--launch_template_configs--overrides--instance_requirements--accelerator_count"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.accelerator_count`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--accelerator_total_memory_mi_b"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.accelerator_total_memory_mi_b`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--baseline_ebs_bandwidth_mbps"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.baseline_ebs_bandwidth_mbps`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.baseline_performance_factors`

Read-Only:

- `cpu` (Attributes) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors--cpu))

<a id="nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors--cpu"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.baseline_performance_factors.cpu`

Read-Only:

- `references` (Attributes List) (see [below for nested schema](#nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors--cpu--references))

<a id="nestedatt--launch_template_configs--overrides--instance_requirements--baseline_performance_factors--cpu--references"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.baseline_performance_factors.cpu.references`

Read-Only:

- `instance_family` (String)




<a id="nestedatt--launch_template_configs--overrides--instance_requirements--memory_gi_b_per_v_cpu"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.memory_gi_b_per_v_cpu`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--memory_mi_b"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.memory_mi_b`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--network_bandwidth_gbps"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.network_bandwidth_gbps`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--network_interface_count"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.network_interface_count`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--total_local_storage_gb"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.total_local_storage_gb`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_configs--overrides--instance_requirements--v_cpu_count"></a>
### Nested Schema for `launch_template_configs.overrides.instance_requirements.v_cpu_count`

Read-Only:

- `max` (Number)
- `min` (Number)



<a id="nestedatt--launch_template_configs--overrides--placement"></a>
### Nested Schema for `launch_template_configs.overrides.placement`

Read-Only:

- `affinity` (String)
- `availability_zone` (String)
- `group_name` (String)
- `host_id` (String)
- `host_resource_group_arn` (String)
- `partition_number` (Number)
- `spread_domain` (String)
- `tenancy` (String)




<a id="nestedatt--on_demand_options"></a>
### Nested Schema for `on_demand_options`

Read-Only:

- `allocation_strategy` (String)
- `capacity_reservation_options` (Attributes) (see [below for nested schema](#nestedatt--on_demand_options--capacity_reservation_options))
- `max_total_price` (String)
- `min_target_capacity` (Number)
- `single_availability_zone` (Boolean)
- `single_instance_type` (Boolean)

<a id="nestedatt--on_demand_options--capacity_reservation_options"></a>
### Nested Schema for `on_demand_options.capacity_reservation_options`

Read-Only:

- `usage_strategy` (String)



<a id="nestedatt--spot_options"></a>
### Nested Schema for `spot_options`

Read-Only:

- `allocation_strategy` (String)
- `instance_interruption_behavior` (String)
- `instance_pools_to_use_count` (Number)
- `maintenance_strategies` (Attributes) (see [below for nested schema](#nestedatt--spot_options--maintenance_strategies))
- `max_total_price` (String)
- `min_target_capacity` (Number)
- `single_availability_zone` (Boolean)
- `single_instance_type` (Boolean)

<a id="nestedatt--spot_options--maintenance_strategies"></a>
### Nested Schema for `spot_options.maintenance_strategies`

Read-Only:

- `capacity_rebalance` (Attributes) (see [below for nested schema](#nestedatt--spot_options--maintenance_strategies--capacity_rebalance))

<a id="nestedatt--spot_options--maintenance_strategies--capacity_rebalance"></a>
### Nested Schema for `spot_options.maintenance_strategies.capacity_rebalance`

Read-Only:

- `replacement_strategy` (String)
- `termination_delay` (Number)




<a id="nestedatt--tag_specifications"></a>
### Nested Schema for `tag_specifications`

Read-Only:

- `resource_type` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tag_specifications--tags))

<a id="nestedatt--tag_specifications--tags"></a>
### Nested Schema for `tag_specifications.tags`

Read-Only:

- `key` (String)
- `value` (String)



<a id="nestedatt--target_capacity_specification"></a>
### Nested Schema for `target_capacity_specification`

Read-Only:

- `default_target_capacity_type` (String)
- `on_demand_target_capacity` (Number)
- `spot_target_capacity` (Number)
- `target_capacity_unit_type` (String)
- `total_target_capacity` (Number)
