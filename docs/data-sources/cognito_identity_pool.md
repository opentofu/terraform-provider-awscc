---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cognito_identity_pool Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Cognito::IdentityPool
---

# awscc_cognito_identity_pool (Data Source)

Data Source schema for AWS::Cognito::IdentityPool



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `allow_classic_flow` (Boolean)
- `allow_unauthenticated_identities` (Boolean)
- `cognito_events` (String)
- `cognito_identity_providers` (Attributes List) (see [below for nested schema](#nestedatt--cognito_identity_providers))
- `cognito_streams` (Attributes) (see [below for nested schema](#nestedatt--cognito_streams))
- `developer_provider_name` (String)
- `identity_pool_id` (String)
- `identity_pool_name` (String)
- `identity_pool_tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--identity_pool_tags))
- `name` (String)
- `open_id_connect_provider_ar_ns` (List of String)
- `push_sync` (Attributes) (see [below for nested schema](#nestedatt--push_sync))
- `saml_provider_ar_ns` (List of String)
- `supported_login_providers` (String)

<a id="nestedatt--cognito_identity_providers"></a>
### Nested Schema for `cognito_identity_providers`

Read-Only:

- `client_id` (String)
- `provider_name` (String)
- `server_side_token_check` (Boolean)


<a id="nestedatt--cognito_streams"></a>
### Nested Schema for `cognito_streams`

Read-Only:

- `role_arn` (String)
- `stream_name` (String)
- `streaming_status` (String)


<a id="nestedatt--identity_pool_tags"></a>
### Nested Schema for `identity_pool_tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.


<a id="nestedatt--push_sync"></a>
### Nested Schema for `push_sync`

Read-Only:

- `application_arns` (List of String)
- `role_arn` (String)
