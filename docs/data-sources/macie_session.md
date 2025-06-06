---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_macie_session Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Macie::Session
---

# awscc_macie_session (Data Source)

Data Source schema for AWS::Macie::Session



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `automated_discovery_status` (String) The status of automated sensitive data discovery for the Macie session.
- `aws_account_id` (String) AWS account ID of customer
- `finding_publishing_frequency` (String) A enumeration value that specifies how frequently finding updates are published.
- `service_role` (String) Service role used by Macie
- `status` (String) A enumeration value that specifies the status of the Macie Session.
