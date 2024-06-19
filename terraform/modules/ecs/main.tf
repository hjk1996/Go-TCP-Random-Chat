resource "aws_ecs_cluster" "main" {
  name = "${var.app_name}-cluster"
}

resource "aws_cloudwatch_log_group" "app_log_group" {
  name = "${var.app_name}-log-group"
}

resource "aws_ecs_task_definition" "app_task_definition" {
  family                   = var.app_name
  cpu                      = 256
  memory                   = 512
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  execution_role_arn = var.app_execution_role_arn
  task_role_arn      = var.app_task_role_arn


  container_definitions = jsonencode([
    {
      name      = var.app_name
      image     = var.app_image
      essential = true
      portMappings = [
        {
          containerPort = var.app_port
          hostPort      = var.app_port
          protocol      = "TCP"
        }
      ]

      environment = concat(
        [
          for k, v in var.app_environment_variables : {
            name  = k
            value = v
          }
        ],
        [
          {
            name  = "REDIS_ADDRESS"
            value = var.redis_endpoint
          }
        ]
      )

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          "awslogs-group"         = aws_cloudwatch_log_group.app_log_group.name
          "awslogs-region"        = var.region
          "awslogs-stream-prefix" = "ecs"
        }
      }

      #   healthCheck = {
      #     command     = ["CMD-SHELL", "curl -f http://localhost:8000/health-check || exit 1"]
      #     interval    = 15
      #     timeout     = 5
      #     retries     = 3
      #     startPeriod = 0
      #   }
    }
  ])
}


resource "aws_ecs_service" "app_ecs_service" {
  name            = "${var.app_name}-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app_task_definition.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = var.app_subnets
    security_groups  = [var.app_task_sg_id]
    assign_public_ip = true
  }
  load_balancer {
    container_port   = var.app_port
    container_name   = var.app_name
    target_group_arn = var.target_group_arn
  }
}