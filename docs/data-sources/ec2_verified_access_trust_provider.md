---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_verified_access_trust_provider Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::VerifiedAccessTrustProvider
---

# awscc_ec2_verified_access_trust_provider (Data Source)

Data Source schema for AWS::EC2::VerifiedAccessTrustProvider



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `creation_time` (String) The creation time.
- `description` (String) A description for the Amazon Web Services Verified Access trust provider.
- `device_options` (Attributes) The options for device identity based trust providers. (see [below for nested schema](#nestedatt--device_options))
- `device_trust_provider_type` (String) The type of device-based trust provider. Possible values: jamf|crowdstrike
- `last_updated_time` (String) The last updated time.
- `oidc_options` (Attributes) The OpenID Connect details for an oidc -type, user-identity based trust provider. (see [below for nested schema](#nestedatt--oidc_options))
- `policy_reference_name` (String) The identifier to be used when working with policy rules.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `trust_provider_type` (String) Type of trust provider. Possible values: user|device
- `user_trust_provider_type` (String) The type of device-based trust provider. Possible values: oidc|iam-identity-center
- `verified_access_trust_provider_id` (String) The ID of the Amazon Web Services Verified Access trust provider.

<a id="nestedatt--device_options"></a>
### Nested Schema for `device_options`

Read-Only:

- `tenant_id` (String) The ID of the tenant application with the device-identity provider.


<a id="nestedatt--oidc_options"></a>
### Nested Schema for `oidc_options`

Read-Only:

- `authorization_endpoint` (String) The OIDC authorization endpoint.
- `client_id` (String) The client identifier.
- `client_secret` (String) The client secret.
- `issuer` (String) The OIDC issuer.
- `scope` (String) OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.
- `token_endpoint` (String) The OIDC token endpoint.
- `user_info_endpoint` (String) The OIDC user info endpoint.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.