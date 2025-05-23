---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_applicationinsights_application Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApplicationInsights::Application
---

# awscc_applicationinsights_application (Data Source)

Data Source schema for AWS::ApplicationInsights::Application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `application_arn` (String) The ARN of the ApplicationInsights application.
- `attach_missing_permission` (Boolean) If set to true, the managed policies for SSM and CW will be attached to the instance roles if they are missing
- `auto_configuration_enabled` (Boolean) If set to true, application will be configured with recommended monitoring configuration.
- `component_monitoring_settings` (Attributes List) The monitoring settings of the components. (see [below for nested schema](#nestedatt--component_monitoring_settings))
- `custom_components` (Attributes List) The custom grouped components. (see [below for nested schema](#nestedatt--custom_components))
- `cwe_monitor_enabled` (Boolean) Indicates whether Application Insights can listen to CloudWatch events for the application resources.
- `grouping_type` (String) The grouping type of the application
- `log_pattern_sets` (Attributes List) The log pattern sets. (see [below for nested schema](#nestedatt--log_pattern_sets))
- `ops_center_enabled` (Boolean) When set to true, creates opsItems for any problems detected on an application.
- `ops_item_sns_topic_arn` (String) The SNS topic provided to Application Insights that is associated to the created opsItem.
- `resource_group_name` (String) The name of the resource group.
- `sns_notification_arn` (String) Application Insights sends notifications to this SNS topic whenever there is a problem update in the associated application.
- `tags` (Attributes List) The tags of Application Insights application. (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--component_monitoring_settings"></a>
### Nested Schema for `component_monitoring_settings`

Read-Only:

- `component_arn` (String) The ARN of the compnonent.
- `component_configuration_mode` (String) The component monitoring configuration mode.
- `component_name` (String) The name of the component.
- `custom_component_configuration` (Attributes) The monitoring configuration of the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration))
- `default_overwrite_component_configuration` (Attributes) The overwritten settings on default component monitoring configuration. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration))
- `tier` (String) The tier of the application component.

<a id="nestedatt--component_monitoring_settings--custom_component_configuration"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration`

Read-Only:

- `configuration_details` (Attributes) The configuration settings (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details))
- `sub_component_type_configurations` (Attributes List) Sub component configurations of the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations))

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--alarm_metrics))
- `alarms` (Attributes List) A list of alarms to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--alarms))
- `ha_cluster_prometheus_exporter` (Attributes) The HA cluster Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--ha_cluster_prometheus_exporter))
- `hana_prometheus_exporter` (Attributes) The HANA DB Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--hana_prometheus_exporter))
- `jmx_prometheus_exporter` (Attributes) The JMX Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--jmx_prometheus_exporter))
- `logs` (Attributes List) A list of logs to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--logs))
- `net_weaver_prometheus_exporter` (Attributes) The NetWeaver Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--net_weaver_prometheus_exporter))
- `processes` (Attributes List) A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--processes))
- `sql_server_prometheus_exporter` (Attributes) The SQL Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--sql_server_prometheus_exporter))
- `windows_events` (Attributes List) A list of Windows Events to log. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--windows_events))

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--alarms"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.alarms`

Read-Only:

- `alarm_name` (String) The name of the CloudWatch alarm to be monitored for the component.
- `severity` (String) Indicates the degree of outage when the alarm goes off.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--ha_cluster_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.ha_cluster_prometheus_exporter`

Read-Only:

- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--hana_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.hana_prometheus_exporter`

Read-Only:

- `agree_to_install_hanadb_client` (Boolean) A flag which indicates agreeing to install SAP HANA DB client.
- `hana_port` (String) The HANA DB port.
- `hana_secret_name` (String) The secret name which manages the HANA DB credentials e.g. {
  "username": "<>",
  "password": "<>"
}.
- `hanasid` (String) HANA DB SID.
- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--jmx_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.jmx_prometheus_exporter`

Read-Only:

- `host_port` (String) Java agent host port
- `jmxurl` (String) JMX service URL.
- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--logs"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.logs`

Read-Only:

- `encoding` (String) The type of encoding of the logs to be monitored.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `log_path` (String) The path of the logs to be monitored.
- `log_type` (String) The log type decides the log patterns against which Application Insights analyzes the log.
- `pattern_set` (String) The name of the log pattern set.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--net_weaver_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.net_weaver_prometheus_exporter`

Read-Only:

- `instance_numbers` (List of String) SAP instance numbers for ASCS, ERS, and App Servers.
- `prometheus_port` (String) Prometheus exporter port.
- `sapsid` (String) SAP NetWeaver SID.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--processes"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.processes`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--processes--alarm_metrics))
- `process_name` (String) The name of the process to be monitored for the component.

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--processes--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.processes.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.



<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--sql_server_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.sql_server_prometheus_exporter`

Read-Only:

- `prometheus_port` (String) Prometheus exporter port.
- `sql_secret_name` (String) Secret name which managers SQL exporter connection. e.g. {"data_source_name": "sqlserver://<USERNAME>:<PASSWORD>@localhost:1433"}


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--configuration_details--windows_events"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.configuration_details.windows_events`

Read-Only:

- `event_levels` (List of String) The levels of event to log.
- `event_name` (String) The type of Windows Events to log.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `pattern_set` (String) The name of the log pattern set.



<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations`

Read-Only:

- `sub_component_configuration_details` (Attributes) The configuration settings of sub components. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details))
- `sub_component_type` (String) The sub component type.

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--alarm_metrics))
- `logs` (Attributes List) A list of logs to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--logs))
- `processes` (Attributes List) A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes))
- `windows_events` (Attributes List) A list of Windows Events to log. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--windows_events))

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--logs"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details.logs`

Read-Only:

- `encoding` (String) The type of encoding of the logs to be monitored.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `log_path` (String) The path of the logs to be monitored.
- `log_type` (String) The log type decides the log patterns against which Application Insights analyzes the log.
- `pattern_set` (String) The name of the log pattern set.


<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details.processes`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes--alarm_metrics))
- `process_name` (String) The name of the process to be monitored for the component.

<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details.processes.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.



<a id="nestedatt--component_monitoring_settings--custom_component_configuration--sub_component_type_configurations--sub_component_configuration_details--windows_events"></a>
### Nested Schema for `component_monitoring_settings.custom_component_configuration.sub_component_type_configurations.sub_component_configuration_details.windows_events`

Read-Only:

- `event_levels` (List of String) The levels of event to log.
- `event_name` (String) The type of Windows Events to log.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `pattern_set` (String) The name of the log pattern set.





<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration`

Read-Only:

- `configuration_details` (Attributes) The configuration settings (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details))
- `sub_component_type_configurations` (Attributes List) Sub component configurations of the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations))

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--alarm_metrics))
- `alarms` (Attributes List) A list of alarms to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--alarms))
- `ha_cluster_prometheus_exporter` (Attributes) The HA cluster Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--ha_cluster_prometheus_exporter))
- `hana_prometheus_exporter` (Attributes) The HANA DB Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--hana_prometheus_exporter))
- `jmx_prometheus_exporter` (Attributes) The JMX Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--jmx_prometheus_exporter))
- `logs` (Attributes List) A list of logs to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--logs))
- `net_weaver_prometheus_exporter` (Attributes) The NetWeaver Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--net_weaver_prometheus_exporter))
- `processes` (Attributes List) A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--processes))
- `sql_server_prometheus_exporter` (Attributes) The SQL Prometheus Exporter settings. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--sql_server_prometheus_exporter))
- `windows_events` (Attributes List) A list of Windows Events to log. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--windows_events))

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--alarms"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.alarms`

Read-Only:

- `alarm_name` (String) The name of the CloudWatch alarm to be monitored for the component.
- `severity` (String) Indicates the degree of outage when the alarm goes off.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--ha_cluster_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.ha_cluster_prometheus_exporter`

Read-Only:

- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--hana_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.hana_prometheus_exporter`

Read-Only:

- `agree_to_install_hanadb_client` (Boolean) A flag which indicates agreeing to install SAP HANA DB client.
- `hana_port` (String) The HANA DB port.
- `hana_secret_name` (String) The secret name which manages the HANA DB credentials e.g. {
  "username": "<>",
  "password": "<>"
}.
- `hanasid` (String) HANA DB SID.
- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--jmx_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.jmx_prometheus_exporter`

Read-Only:

- `host_port` (String) Java agent host port
- `jmxurl` (String) JMX service URL.
- `prometheus_port` (String) Prometheus exporter port.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--logs"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.logs`

Read-Only:

- `encoding` (String) The type of encoding of the logs to be monitored.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `log_path` (String) The path of the logs to be monitored.
- `log_type` (String) The log type decides the log patterns against which Application Insights analyzes the log.
- `pattern_set` (String) The name of the log pattern set.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--net_weaver_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.net_weaver_prometheus_exporter`

Read-Only:

- `instance_numbers` (List of String) SAP instance numbers for ASCS, ERS, and App Servers.
- `prometheus_port` (String) Prometheus exporter port.
- `sapsid` (String) SAP NetWeaver SID.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--processes"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.processes`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--processes--alarm_metrics))
- `process_name` (String) The name of the process to be monitored for the component.

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--processes--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.processes.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.



<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--sql_server_prometheus_exporter"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.sql_server_prometheus_exporter`

Read-Only:

- `prometheus_port` (String) Prometheus exporter port.
- `sql_secret_name` (String) Secret name which managers SQL exporter connection. e.g. {"data_source_name": "sqlserver://<USERNAME>:<PASSWORD>@localhost:1433"}


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--configuration_details--windows_events"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.configuration_details.windows_events`

Read-Only:

- `event_levels` (List of String) The levels of event to log.
- `event_name` (String) The type of Windows Events to log.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `pattern_set` (String) The name of the log pattern set.



<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations`

Read-Only:

- `sub_component_configuration_details` (Attributes) The configuration settings of sub components. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details))
- `sub_component_type` (String) The sub component type.

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--alarm_metrics))
- `logs` (Attributes List) A list of logs to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--logs))
- `processes` (Attributes List) A list of processes to monitor for the component. Only Windows EC2 instances can have a processes section. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes))
- `windows_events` (Attributes List) A list of Windows Events to log. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--windows_events))

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--logs"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details.logs`

Read-Only:

- `encoding` (String) The type of encoding of the logs to be monitored.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `log_path` (String) The path of the logs to be monitored.
- `log_type` (String) The log type decides the log patterns against which Application Insights analyzes the log.
- `pattern_set` (String) The name of the log pattern set.


<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details.processes`

Read-Only:

- `alarm_metrics` (Attributes List) A list of metrics to monitor for the component. (see [below for nested schema](#nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes--alarm_metrics))
- `process_name` (String) The name of the process to be monitored for the component.

<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--processes--alarm_metrics"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details.processes.alarm_metrics`

Read-Only:

- `alarm_metric_name` (String) The name of the metric to be monitored for the component.



<a id="nestedatt--component_monitoring_settings--default_overwrite_component_configuration--sub_component_type_configurations--sub_component_configuration_details--windows_events"></a>
### Nested Schema for `component_monitoring_settings.default_overwrite_component_configuration.sub_component_type_configurations.sub_component_configuration_details.windows_events`

Read-Only:

- `event_levels` (List of String) The levels of event to log.
- `event_name` (String) The type of Windows Events to log.
- `log_group_name` (String) The CloudWatch log group name to be associated to the monitored log.
- `pattern_set` (String) The name of the log pattern set.






<a id="nestedatt--custom_components"></a>
### Nested Schema for `custom_components`

Read-Only:

- `component_name` (String) The name of the component.
- `resource_list` (List of String) The list of resource ARNs that belong to the component.


<a id="nestedatt--log_pattern_sets"></a>
### Nested Schema for `log_pattern_sets`

Read-Only:

- `log_patterns` (Attributes List) The log patterns of a set. (see [below for nested schema](#nestedatt--log_pattern_sets--log_patterns))
- `pattern_set_name` (String) The name of the log pattern set.

<a id="nestedatt--log_pattern_sets--log_patterns"></a>
### Nested Schema for `log_pattern_sets.log_patterns`

Read-Only:

- `pattern` (String) The log pattern.
- `pattern_name` (String) The name of the log pattern.
- `rank` (Number) Rank of the log pattern.



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
