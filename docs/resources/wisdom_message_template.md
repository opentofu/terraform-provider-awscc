---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_wisdom_message_template Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Wisdom::MessageTemplate Resource Type
---

# awscc_wisdom_message_template (Resource)

Definition of AWS::Wisdom::MessageTemplate Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `channel_subtype` (String) The channel subtype this message template applies to.
- `content` (Attributes) The content of the message template. (see [below for nested schema](#nestedatt--content))
- `knowledge_base_arn` (String) The Amazon Resource Name (ARN) of the knowledge base to which the message template belongs.
- `name` (String) The name of the message template.

### Optional

- `default_attributes` (Attributes) An object that specifies the default values to use for variables in the message template. This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable. (see [below for nested schema](#nestedatt--default_attributes))
- `description` (String) The description of the message template.
- `grouping_configuration` (Attributes) The configuration information of the user groups that the message template is accessible to. (see [below for nested schema](#nestedatt--grouping_configuration))
- `language` (String) The language code value for the language in which the message template is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW
- `tags` (Attributes Set) The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `message_template_arn` (String) The Amazon Resource Name (ARN) of the message template.
- `message_template_content_sha_256` (String) The content SHA256 of the message template.
- `message_template_id` (String) The unique identifier of the message template.

<a id="nestedatt--content"></a>
### Nested Schema for `content`

Optional:

- `email_message_template_content` (Attributes) The content of message template that applies to email channel subtype. (see [below for nested schema](#nestedatt--content--email_message_template_content))
- `sms_message_template_content` (Attributes) The content of message template that applies to SMS channel subtype. (see [below for nested schema](#nestedatt--content--sms_message_template_content))

<a id="nestedatt--content--email_message_template_content"></a>
### Nested Schema for `content.email_message_template_content`

Optional:

- `body` (Attributes) The body to use in email messages. (see [below for nested schema](#nestedatt--content--email_message_template_content--body))
- `headers` (Attributes List) The email headers to include in email messages. (see [below for nested schema](#nestedatt--content--email_message_template_content--headers))
- `subject` (String) The subject line, or title, to use in email messages.

<a id="nestedatt--content--email_message_template_content--body"></a>
### Nested Schema for `content.email_message_template_content.body`

Optional:

- `html` (Attributes) The message body, in HTML format, to use in email messages that are based on the message template. We recommend using HTML format for email clients that render HTML content. You can include links, formatted text, and more in an HTML message. (see [below for nested schema](#nestedatt--content--email_message_template_content--body--html))
- `plain_text` (Attributes) The message body, in plain text format, to use in email messages that are based on the message template. We recommend using plain text format for email clients that don't render HTML content and clients that are connected to high-latency networks, such as mobile devices. (see [below for nested schema](#nestedatt--content--email_message_template_content--body--plain_text))

<a id="nestedatt--content--email_message_template_content--body--html"></a>
### Nested Schema for `content.email_message_template_content.body.html`

Optional:

- `content` (String)


<a id="nestedatt--content--email_message_template_content--body--plain_text"></a>
### Nested Schema for `content.email_message_template_content.body.plain_text`

Optional:

- `content` (String)



<a id="nestedatt--content--email_message_template_content--headers"></a>
### Nested Schema for `content.email_message_template_content.headers`

Optional:

- `name` (String) The name of the email header.
- `value` (String) The value of the email header.



<a id="nestedatt--content--sms_message_template_content"></a>
### Nested Schema for `content.sms_message_template_content`

Optional:

- `body` (Attributes) The body to use in SMS messages. (see [below for nested schema](#nestedatt--content--sms_message_template_content--body))

<a id="nestedatt--content--sms_message_template_content--body"></a>
### Nested Schema for `content.sms_message_template_content.body`

Optional:

- `plain_text` (Attributes) The container of message template body. (see [below for nested schema](#nestedatt--content--sms_message_template_content--body--plain_text))

<a id="nestedatt--content--sms_message_template_content--body--plain_text"></a>
### Nested Schema for `content.sms_message_template_content.body.plain_text`

Optional:

- `content` (String)





<a id="nestedatt--default_attributes"></a>
### Nested Schema for `default_attributes`

Optional:

- `agent_attributes` (Attributes) The agent attributes that are used with the message template. (see [below for nested schema](#nestedatt--default_attributes--agent_attributes))
- `custom_attributes` (Map of String) The custom attributes that are used with the message template.
- `customer_profile_attributes` (Attributes) The customer profile attributes that are used with the message template. (see [below for nested schema](#nestedatt--default_attributes--customer_profile_attributes))
- `system_attributes` (Attributes) The system attributes that are used with the message template. (see [below for nested schema](#nestedatt--default_attributes--system_attributes))

<a id="nestedatt--default_attributes--agent_attributes"></a>
### Nested Schema for `default_attributes.agent_attributes`

Optional:

- `first_name` (String) The agent?s first name as entered in their Amazon Connect user account.
- `last_name` (String) The agent?s last name as entered in their Amazon Connect user account.


<a id="nestedatt--default_attributes--customer_profile_attributes"></a>
### Nested Schema for `default_attributes.customer_profile_attributes`

Optional:

- `account_number` (String) A unique account number that you have given to the customer.
- `additional_information` (String) Any additional information relevant to the customer's profile.
- `address_1` (String) The first line of a customer address.
- `address_2` (String) The second line of a customer address.
- `address_3` (String) The third line of a customer address.
- `address_4` (String) The fourth line of a customer address.
- `billing_address_1` (String) The first line of a customer?s billing address.
- `billing_address_2` (String) The second line of a customer?s billing address.
- `billing_address_3` (String) The third line of a customer?s billing address.
- `billing_address_4` (String) The fourth line of a customer?s billing address.
- `billing_city` (String) The city of a customer?s billing address.
- `billing_country` (String) The country of a customer?s billing address.
- `billing_county` (String) The county of a customer?s billing address.
- `billing_postal_code` (String) The postal code of a customer?s billing address.
- `billing_province` (String) The province of a customer?s billing address.
- `billing_state` (String) The state of a customer?s billing address.
- `birth_date` (String) The customer's birth date.
- `business_email_address` (String) The customer's business email address.
- `business_name` (String) The name of the customer's business.
- `business_phone_number` (String) The customer's business phone number.
- `city` (String) The city in which a customer lives.
- `country` (String) The country in which a customer lives.
- `county` (String) The county in which a customer lives.
- `custom` (Map of String) The custom attributes that are used with the message template.
- `email_address` (String) The customer's email address, which has not been specified as a personal or business address.
- `first_name` (String) The customer's first name.
- `gender` (String) The customer's gender.
- `home_phone_number` (String) The customer's home phone number.
- `last_name` (String) The customer's last name.
- `mailing_address_1` (String) The first line of a customer?s mailing address.
- `mailing_address_2` (String) The second line of a customer?s mailing address.
- `mailing_address_3` (String) The third line of a customer?s mailing address.
- `mailing_address_4` (String) The fourth line of a customer?s mailing address.
- `mailing_city` (String) The city of a customer?s mailing address.
- `mailing_country` (String) The country of a customer?s mailing address.
- `mailing_county` (String) The county of a customer?s mailing address.
- `mailing_postal_code` (String) The postal code of a customer?s mailing address
- `mailing_province` (String) The province of a customer?s mailing address.
- `mailing_state` (String) The state of a customer?s mailing address.
- `middle_name` (String) The customer's middle name.
- `mobile_phone_number` (String) The customer's mobile phone number.
- `party_type` (String) The customer's party type.
- `phone_number` (String) The customer's phone number, which has not been specified as a mobile, home, or business number.
- `postal_code` (String) The postal code of a customer address.
- `profile_arn` (String) The ARN of a customer profile.
- `profile_id` (String) The unique identifier of a customer profile.
- `province` (String) The province in which a customer lives.
- `shipping_address_1` (String) The first line of a customer?s shipping address.
- `shipping_address_2` (String) The second line of a customer?s shipping address.
- `shipping_address_3` (String) The third line of a customer?s shipping address.
- `shipping_address_4` (String) The fourth line of a customer?s shipping address
- `shipping_city` (String) The city of a customer?s shipping address.
- `shipping_country` (String) The country of a customer?s shipping address.
- `shipping_county` (String) The county of a customer?s shipping address.
- `shipping_postal_code` (String) The postal code of a customer?s shipping address.
- `shipping_province` (String) The province of a customer?s shipping address.
- `shipping_state` (String) The state of a customer?s shipping address.
- `state` (String) The state in which a customer lives.


<a id="nestedatt--default_attributes--system_attributes"></a>
### Nested Schema for `default_attributes.system_attributes`

Optional:

- `customer_endpoint` (Attributes) The CustomerEndpoint attribute. (see [below for nested schema](#nestedatt--default_attributes--system_attributes--customer_endpoint))
- `name` (String) The name of the task.
- `system_endpoint` (Attributes) The SystemEndpoint attribute. (see [below for nested schema](#nestedatt--default_attributes--system_attributes--system_endpoint))

<a id="nestedatt--default_attributes--system_attributes--customer_endpoint"></a>
### Nested Schema for `default_attributes.system_attributes.customer_endpoint`

Optional:

- `address` (String) The customer's phone number if used with customerEndpoint, or the number the customer dialed to call your contact center if used with systemEndpoint.


<a id="nestedatt--default_attributes--system_attributes--system_endpoint"></a>
### Nested Schema for `default_attributes.system_attributes.system_endpoint`

Optional:

- `address` (String) The customer's phone number if used with customerEndpoint, or the number the customer dialed to call your contact center if used with systemEndpoint.




<a id="nestedatt--grouping_configuration"></a>
### Nested Schema for `grouping_configuration`

Optional:

- `criteria` (String) The criteria used for grouping Amazon Q in Connect users.
- `values` (List of String) The list of values that define different groups of Amazon Q in Connect users.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_wisdom_message_template.example "message_template_arn"
```