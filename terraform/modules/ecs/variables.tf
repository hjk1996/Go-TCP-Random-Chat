variable "region" {
  type = string
}

variable "app_name" {
  type = string
}

variable "app_execution_role_arn" {
  type = string
}

variable "app_task_role_arn" {
  type = string
}

variable "app_port" {
  type = number
}

variable "app_subnets" {
  type = list(string)
}

variable "app_task_sg_id" {
  type = string
}

variable "app_image" {
  type = string
}


variable "app_environment_variables" {
  type = map(string)
}


variable "redis_endpoint" {
  type = string

}

variable "target_group_arn" {
  type = string
}


variable "min_capacity" {
  type = number
}