---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_integration_association Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::IntegrationAssociation
---

# awscc_connect_integration_association (Data Source)

Data Source schema for AWS::Connect::IntegrationAssociation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `instance_id` (String) Amazon Connect instance identifier
- `integration_arn` (String) ARN of Integration being associated with the instance
- `integration_association_id` (String) Identifier of the association with Connect Instance
- `integration_type` (String) Specifies the integration type to be associated with the instance

