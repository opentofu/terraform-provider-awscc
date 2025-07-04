
---
page_title: "awscc_ecs_task_set Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
---

# awscc_ecs_task_set (Resource)

Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.

## Example Usage

### Deploy ECS Task Set with Load Balancer

Creates an ECS Task Set running Nginx containers in Fargate mode with Application Load Balancer integration, complete with networking, security, and IAM configurations.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# VPC and Networking Resources
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "ECS Task Set Demo VPC"
  }
}

resource "aws_subnet" "public1" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "ECS Task Set Public Subnet 1"
  }
}

resource "aws_subnet" "public2" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = {
    Name = "ECS Task Set Public Subnet 2"
  }
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "ECS Task Set IGW"
  }
}

resource "aws_security_group" "ecs" {
  name        = "ecs-task-set-sg"
  description = "Security group for ECS tasks"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Load Balancer Resources
resource "aws_lb" "main" {
  name               = "ecs-task-set-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.ecs.id]
  subnets            = [aws_subnet.public1.id, aws_subnet.public2.id]
}

resource "aws_lb_target_group" "main" {
  name        = "ecs-task-set-tg"
  port        = 80
  protocol    = "HTTP"
  vpc_id      = aws_vpc.main.id
  target_type = "ip"
}

resource "aws_lb_listener" "main" {
  load_balancer_arn = aws_lb.main.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.main.arn
  }
}

# IAM Resources
resource "awscc_iam_role" "ecs_task_execution" {
  role_name = "ecs-task-set-execution-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"]

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_iam_role" "ecs_task" {
  role_name = "ecs-task-set-role"
  assume_role_policy_document = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# ECS Resources
resource "awscc_ecs_cluster" "main" {
  cluster_name = "task-set-demo"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_task_definition" "app" {
  family = "task-set-app"

  container_definitions = [
    {
      name      = "nginx"
      image     = "nginx:latest"
      cpu       = 256
      memory    = 512
      essential = true
      port_mappings = [
        {
          container_port = 80
          protocol       = "tcp"
        }
      ]
    }
  ]

  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "512"
  memory                   = "1024"
  execution_role_arn       = awscc_iam_role.ecs_task_execution.arn
  task_role_arn            = awscc_iam_role.ecs_task.arn

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_service" "app" {
  service_name  = "task-set-service"
  cluster       = awscc_ecs_cluster.main.id
  desired_count = 2

  deployment_configuration = {
    deployment_circuit_breaker = {
      enable   = true
      rollback = true
    }
    maximum_percent         = 200
    minimum_healthy_percent = 100
  }

  load_balancers = [
    {
      container_name   = "nginx"
      container_port   = 80
      target_group_arn = aws_lb_target_group.main.arn
    }
  ]

  network_configuration = {
    aws_vpc_configuration = {
      assign_public_ip = true
      security_groups  = [aws_security_group.ecs.id]
      subnets          = [aws_subnet.public1.id, aws_subnet.public2.id]
    }
  }

  task_definition = awscc_ecs_task_definition.app.id

  scheduling_strategy = "REPLICA"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

resource "awscc_ecs_task_set" "app_v1" {
  service     = awscc_ecs_service.app.id
  cluster     = awscc_ecs_cluster.main.id
  external_id = "v1"
  scale = {
    unit  = "PERCENT"
    value = 100
  }
  task_definition = awscc_ecs_task_definition.app.id

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster` (String) The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
- `service` (String) The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
- `task_definition` (String) The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.

### Optional

- `capacity_provider_strategy` (Attributes List) (see [below for nested schema](#nestedatt--capacity_provider_strategy))
- `external_id` (String) An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value.
- `launch_type` (String) The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide.
- `load_balancers` (Attributes List) (see [below for nested schema](#nestedatt--load_balancers))
- `network_configuration` (Attributes) An object representing the network configuration for a task or service. (see [below for nested schema](#nestedatt--network_configuration))
- `platform_version` (String) The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.
- `scale` (Attributes) A floating-point percentage of the desired number of tasks to place and keep running in the task set. (see [below for nested schema](#nestedatt--scale))
- `service_registries` (Attributes List) The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html. (see [below for nested schema](#nestedatt--service_registries))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `task_set_id` (String) The ID of the task set.

<a id="nestedatt--capacity_provider_strategy"></a>
### Nested Schema for `capacity_provider_strategy`

Optional:

- `base` (Number)
- `capacity_provider` (String)
- `weight` (Number)


<a id="nestedatt--load_balancers"></a>
### Nested Schema for `load_balancers`

Optional:

- `container_name` (String) The name of the container (as it appears in a container definition) to associate with the load balancer.
- `container_port` (Number) The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.
- `target_group_arn` (String) The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.


<a id="nestedatt--network_configuration"></a>
### Nested Schema for `network_configuration`

Optional:

- `aws_vpc_configuration` (Attributes) The VPC subnets and security groups associated with a task. All specified subnets and security groups must be from the same VPC. (see [below for nested schema](#nestedatt--network_configuration--aws_vpc_configuration))

<a id="nestedatt--network_configuration--aws_vpc_configuration"></a>
### Nested Schema for `network_configuration.aws_vpc_configuration`

Optional:

- `assign_public_ip` (String) Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.
- `security_groups` (List of String) The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.
- `subnets` (List of String) The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.



<a id="nestedatt--scale"></a>
### Nested Schema for `scale`

Optional:

- `unit` (String) The unit of measure for the scale value.
- `value` (Number) The value, specified as a percent total of a service's desiredCount, to scale the task set. Accepted values are numbers between 0 and 100.


<a id="nestedatt--service_registries"></a>
### Nested Schema for `service_registries`

Optional:

- `container_name` (String) The container name value, already specified in the task definition, to be used for your service discovery service. If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.
- `container_port` (Number) The port value, already specified in the task definition, to be used for your service discovery service. If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.
- `port` (Number) The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.
- `registry_arn` (String) The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html


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
  to = awscc_ecs_task_set.example
  id = "cluster|service|id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ecs_task_set.example "cluster|service|id"
```
