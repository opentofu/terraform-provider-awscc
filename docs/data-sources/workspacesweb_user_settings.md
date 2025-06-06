---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_workspacesweb_user_settings Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::WorkSpacesWeb::UserSettings
---

# awscc_workspacesweb_user_settings (Data Source)

Data Source schema for AWS::WorkSpacesWeb::UserSettings



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `additional_encryption_context` (Map of String)
- `associated_portal_arns` (List of String)
- `cookie_synchronization_configuration` (Attributes) (see [below for nested schema](#nestedatt--cookie_synchronization_configuration))
- `copy_allowed` (String)
- `customer_managed_key` (String)
- `deep_link_allowed` (String)
- `disconnect_timeout_in_minutes` (Number)
- `download_allowed` (String)
- `idle_disconnect_timeout_in_minutes` (Number)
- `paste_allowed` (String)
- `print_allowed` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `toolbar_configuration` (Attributes) (see [below for nested schema](#nestedatt--toolbar_configuration))
- `upload_allowed` (String)
- `user_settings_arn` (String)

<a id="nestedatt--cookie_synchronization_configuration"></a>
### Nested Schema for `cookie_synchronization_configuration`

Read-Only:

- `allowlist` (Attributes List) (see [below for nested schema](#nestedatt--cookie_synchronization_configuration--allowlist))
- `blocklist` (Attributes List) (see [below for nested schema](#nestedatt--cookie_synchronization_configuration--blocklist))

<a id="nestedatt--cookie_synchronization_configuration--allowlist"></a>
### Nested Schema for `cookie_synchronization_configuration.allowlist`

Read-Only:

- `domain` (String)
- `name` (String)
- `path` (String)


<a id="nestedatt--cookie_synchronization_configuration--blocklist"></a>
### Nested Schema for `cookie_synchronization_configuration.blocklist`

Read-Only:

- `domain` (String)
- `name` (String)
- `path` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)


<a id="nestedatt--toolbar_configuration"></a>
### Nested Schema for `toolbar_configuration`

Read-Only:

- `hidden_toolbar_items` (List of String)
- `max_display_resolution` (String)
- `toolbar_type` (String)
- `visual_mode` (String)
