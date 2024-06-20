module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "${var.app_name}-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["${var.region}a", "${var.region}b", "${var.region}c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}


resource "aws_security_group" "app_task_sg" {
  name   = "${var.app_name}-task-sg"
  vpc_id = module.vpc.vpc_id
  ingress {
    protocol    = "TCP"
    from_port   = var.app_port
    to_port     = var.app_port
    cidr_blocks = module.vpc.public_subnets_cidr_blocks
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}


resource "aws_security_group" "app_redis_sg" {
  name   = "${var.app_name}-redis-sg"
  vpc_id = module.vpc.vpc_id
  ingress {
    protocol    = "TCP"
    from_port   = 6379
    to_port     = 6379
    cidr_blocks = module.vpc.public_subnets_cidr_blocks
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }


}



///lb

resource "aws_security_group" "app_lb_sg" {
  name   = "${var.app_name}-lb-sg"
  vpc_id = module.vpc.vpc_id
  ingress {
    protocol    = "TCP"
    from_port   = var.app_port
    to_port     = var.app_port
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }


}

// lb

resource "aws_lb" "app_lb" {
  name               = "${var.app_name}-lb"
  internal           = false
  load_balancer_type = "network"
  security_groups    = [aws_security_group.app_lb_sg.id]
  subnets            = module.vpc.public_subnets

}

resource "aws_lb_target_group" "app_lb_tg" {
  name     = "${var.app_name}-lb-tg"
  port     = var.app_port
  protocol = "TCP"
  vpc_id   = module.vpc.vpc_id
  // ecs fargate로 container 배포할 땐 awsvpc 모드로 배포되므로 target_type을 ip로 설정해야한다
  target_type = "ip"



  stickiness {
    type            = "source_ip"
    enabled         = true
    cookie_duration = 86400
  }

  lifecycle {
    create_before_destroy = true
  }

}


resource "aws_lb_listener" "app_lb_listener" {
  load_balancer_arn = aws_lb.app_lb.arn
  port              = var.app_port
  protocol          = "TCP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.app_lb_tg.arn
  }

  lifecycle {
    create_before_destroy = true
  }

  depends_on = [aws_lb_target_group.app_lb_tg]

}

