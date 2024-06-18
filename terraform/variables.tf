
variable "region" {
  type    = string
  default = "ap-northeast-2"
}

variable "aws_access_key" {
  type = string
}

variable "aws_secret_key" {
  type = string
}

variable "app_name" {
  type = string
}

variable "app_port" {
  type    = number
  default = 8888
}


variable "app_environment_variables" {
  type = map(string)
}

variable "app_image" {
  type        = string
  description = "docker image path"
}


variable "redis_num_nodes" {
  type = number
}

variable "redis_node_type" {
  type = string
}