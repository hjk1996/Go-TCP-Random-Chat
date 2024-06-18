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

///lb

resource "aws_lb_target_group" "app_tg" {
  name     = "${var.app_name}-ecs-tg"
  port     = var.app_port
  protocol = "TCP"
  vpc_id   = module.vpc.vpc_id


}


# resource "aws_lb" "name" {

# }