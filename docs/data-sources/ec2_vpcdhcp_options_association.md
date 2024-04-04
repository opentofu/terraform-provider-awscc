---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_vpcdhcp_options_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::VPCDHCPOptionsAssociation
---

# awscc_ec2_vpcdhcp_options_association (Data Source)

Data Source schema for AWS::EC2::VPCDHCPOptionsAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `dhcp_options_id` (String) The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
- `vpc_id` (String) The ID of the VPC.