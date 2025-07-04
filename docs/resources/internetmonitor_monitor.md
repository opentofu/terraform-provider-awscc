
---
page_title: "awscc_internetmonitor_monitor Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
---

# awscc_internetmonitor_monitor (Resource)

Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application

## Example Usage

### Configure Internet Monitor with S3 Logging

Creates an Amazon CloudWatch Internet Monitor that monitors city networks and traffic patterns, with health event configurations and measurements logged to S3, configured to monitor 100 city networks and 20% of traffic with availability and performance thresholds set to 95%.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS account ID and region
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Create an S3 bucket for the monitor logs
resource "awscc_s3_bucket" "monitor_logs" {
  bucket_name = "internetmonitor-logs-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create a bucket policy to allow Internet Monitor to write logs
# Note: We still need aws_iam_policy_document data source as it's a helper for policy generation
data "aws_iam_policy_document" "monitor_bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["internetmonitor.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = [
      "arn:aws:s3:::${awscc_s3_bucket.monitor_logs.bucket_name}/*"
    ]
  }
}

resource "awscc_s3_bucket_policy" "monitor_logs" {
  bucket          = awscc_s3_bucket.monitor_logs.bucket_name
  policy_document = data.aws_iam_policy_document.monitor_bucket_policy.json
}

# Create the Internet Monitor monitor
resource "awscc_internetmonitor_monitor" "example" {
  monitor_name = "example-monitor"

  # Monitor 100 city networks
  max_city_networks_to_monitor = 100

  # Monitor 20% of traffic
  traffic_percentage_to_monitor = 20

  # Define health events configuration
  health_events_config = {
    availability_score_threshold = 95
    performance_score_threshold  = 95

    availability_local_health_events_config = {
      health_score_threshold = 90
      min_traffic_impact     = 0.2
      status                 = "ENABLED"
    }

    performance_local_health_events_config = {
      health_score_threshold = 90
      min_traffic_impact     = 0.2
      status                 = "ENABLED"
    }
  }

  # Configure log delivery to S3
  internet_measurements_log_delivery = {
    s3_config = {
      bucket_name         = awscc_s3_bucket.monitor_logs.bucket_name
      bucket_prefix       = "internet-monitor-logs"
      log_delivery_status = "ENABLED"
    }
  }

  # Add example resources to monitor
  # Only setting monitor name and log delivery for example
  resources = []

  # Add tags
  tags = [{
    key   = "Environment"
    value = "Production"
    }, {
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `monitor_name` (String)

### Optional

- `health_events_config` (Attributes) (see [below for nested schema](#nestedatt--health_events_config))
- `include_linked_accounts` (Boolean)
- `internet_measurements_log_delivery` (Attributes) (see [below for nested schema](#nestedatt--internet_measurements_log_delivery))
- `linked_account_id` (String)
- `max_city_networks_to_monitor` (Number)
- `resources` (List of String)
- `resources_to_add` (List of String)
- `resources_to_remove` (List of String)
- `status` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `traffic_percentage_to_monitor` (Number)

### Read-Only

- `created_at` (String) The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)
- `id` (String) Uniquely identifies the resource.
- `modified_at` (String) The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)
- `monitor_arn` (String)
- `processing_status` (String)
- `processing_status_info` (String)

<a id="nestedatt--health_events_config"></a>
### Nested Schema for `health_events_config`

Optional:

- `availability_local_health_events_config` (Attributes) (see [below for nested schema](#nestedatt--health_events_config--availability_local_health_events_config))
- `availability_score_threshold` (Number)
- `performance_local_health_events_config` (Attributes) (see [below for nested schema](#nestedatt--health_events_config--performance_local_health_events_config))
- `performance_score_threshold` (Number)

<a id="nestedatt--health_events_config--availability_local_health_events_config"></a>
### Nested Schema for `health_events_config.availability_local_health_events_config`

Optional:

- `health_score_threshold` (Number)
- `min_traffic_impact` (Number)
- `status` (String)


<a id="nestedatt--health_events_config--performance_local_health_events_config"></a>
### Nested Schema for `health_events_config.performance_local_health_events_config`

Optional:

- `health_score_threshold` (Number)
- `min_traffic_impact` (Number)
- `status` (String)



<a id="nestedatt--internet_measurements_log_delivery"></a>
### Nested Schema for `internet_measurements_log_delivery`

Optional:

- `s3_config` (Attributes) (see [below for nested schema](#nestedatt--internet_measurements_log_delivery--s3_config))

<a id="nestedatt--internet_measurements_log_delivery--s3_config"></a>
### Nested Schema for `internet_measurements_log_delivery.s3_config`

Optional:

- `bucket_name` (String)
- `bucket_prefix` (String)
- `log_delivery_status` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_internetmonitor_monitor.example
  id = "monitor_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_internetmonitor_monitor.example "monitor_name"
```
