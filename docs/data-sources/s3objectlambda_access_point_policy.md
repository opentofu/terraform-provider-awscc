---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_s3objectlambda_access_point_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::S3ObjectLambda::AccessPointPolicy
---

# awscc_s3objectlambda_access_point_policy (Data Source)

Data Source schema for AWS::S3ObjectLambda::AccessPointPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **object_lambda_access_point** (String) The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies.
- **policy_document** (Map of String) A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide.

