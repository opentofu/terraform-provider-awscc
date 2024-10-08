---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_resiliencehub_resiliency_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type Definition for Resiliency Policy.
---

# awscc_resiliencehub_resiliency_policy (Resource)

Resource Type Definition for Resiliency Policy.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `policy` (Attributes) (see [below for nested schema](#nestedatt--policy))
- `policy_name` (String) Name of Resiliency Policy.
- `tier` (String) Resiliency Policy Tier.

### Optional

- `data_location_constraint` (String) Data Location Constraint of the Policy.
- `policy_description` (String) Description of Resiliency Policy.
- `tags` (Map of String)

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `policy_arn` (String) Amazon Resource Name (ARN) of the Resiliency Policy.

<a id="nestedatt--policy"></a>
### Nested Schema for `policy`

Required:

- `az` (Attributes) Failure Policy. (see [below for nested schema](#nestedatt--policy--az))
- `hardware` (Attributes) Failure Policy. (see [below for nested schema](#nestedatt--policy--hardware))
- `software` (Attributes) Failure Policy. (see [below for nested schema](#nestedatt--policy--software))

Optional:

- `region` (Attributes) Failure Policy. (see [below for nested schema](#nestedatt--policy--region))

<a id="nestedatt--policy--az"></a>
### Nested Schema for `policy.az`

Required:

- `rpo_in_secs` (Number) RPO in seconds.
- `rto_in_secs` (Number) RTO in seconds.


<a id="nestedatt--policy--hardware"></a>
### Nested Schema for `policy.hardware`

Required:

- `rpo_in_secs` (Number) RPO in seconds.
- `rto_in_secs` (Number) RTO in seconds.


<a id="nestedatt--policy--software"></a>
### Nested Schema for `policy.software`

Required:

- `rpo_in_secs` (Number) RPO in seconds.
- `rto_in_secs` (Number) RTO in seconds.


<a id="nestedatt--policy--region"></a>
### Nested Schema for `policy.region`

Optional:

- `rpo_in_secs` (Number) RPO in seconds.
- `rto_in_secs` (Number) RTO in seconds.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_resiliencehub_resiliency_policy.example "policy_arn"
```
