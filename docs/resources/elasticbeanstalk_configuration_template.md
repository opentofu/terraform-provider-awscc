---
page_title: "awscc_elasticbeanstalk_configuration_template Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate
---

# awscc_elasticbeanstalk_configuration_template (Resource)

Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate

## Example Usage

### Basic usage with Python platform

In this example, we create an AWS Elastic Beanstalk configuration template using the Python platform, a t3 medium instance and an Immutable deployment policy. This example assumes that you have deployed an environment with high availability and that you are using a dedicated Application Load Balancer. Please refer to [Python platform history](https://docs.aws.amazon.com/elasticbeanstalk/latest/platforms/platform-history-python.html) and use the current platform version and solution stack name. Replace the `testRoleEC2` with your existing [instance profile](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-instanceprofile.html).
```terraform
resource "awscc_elasticbeanstalk_application" "example" {
  application_name = "example"
  description      = "example"
}
resource "awscc_elasticbeanstalk_configuration_template" "example" {
  application_name    = awscc_elasticbeanstalk_application.example.application_name
  description         = "My sample configuration template"
  solution_stack_name = "64bit Amazon Linux 2023 v4.1.2 running Python 3.11"
  option_settings = [{
    namespace   = "aws:autoscaling:launchconfiguration"
    option_name = "IamInstanceProfile"
    value       = "testRoleEC2"
    },
    {
      namespace   = "aws:autoscaling:launchconfiguration"
      option_name = "InstanceType"
      value       = "t3.large"
    },
    {
      namespace   = "aws:elasticbeanstalk:command"
      option_name = "DeploymentPolicy"
      value       = "Immutable"
    },
    {
      namespace   = "aws:elasticbeanstalk:environment"
      option_name = "LoadBalancerType"
      value       = "application"
    },
    {
      namespace   = "aws:elbv2:loadbalancer"
      option_name = "AccessLogsS3Enabled"
      value       = "false"
    },
    {
      namespace   = "aws:elasticbeanstalk:environment"
      option_name = "LoadBalancerIsShared"
      value       = "false"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_name` (String) The name of the Elastic Beanstalk application to associate with this configuration template.

### Optional

- `description` (String) An optional description for this configuration.
- `environment_id` (String) The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration.
- `option_settings` (Attributes List) Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. (see [below for nested schema](#nestedatt--option_settings))
- `platform_arn` (String) The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide.
- `solution_stack_name` (String) The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide.

 You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration.

 Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks.
- `source_configuration` (Attributes) An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.

Values specified in OptionSettings override any values obtained from the SourceConfiguration.

You must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.

Constraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name. (see [below for nested schema](#nestedatt--source_configuration))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `template_name` (String) The name of the configuration template

<a id="nestedatt--option_settings"></a>
### Nested Schema for `option_settings`

Optional:

- `namespace` (String) A unique namespace that identifies the option's associated AWS resource.
- `option_name` (String) The name of the configuration option.
- `resource_name` (String) A unique resource name for the option setting. Use it for a time–based scaling configuration option.
- `value` (String) The current value for the configuration option.


<a id="nestedatt--source_configuration"></a>
### Nested Schema for `source_configuration`

Optional:

- `application_name` (String) The name of the application associated with the configuration.
- `template_name` (String) The name of the configuration template.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_elasticbeanstalk_configuration_template.example
  id = "application_name|template_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_elasticbeanstalk_configuration_template.example "application_name|template_name"
```