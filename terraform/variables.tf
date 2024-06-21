
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
  default = "hjk1996/go-tcp-random-chat:latest"
}


variable "redis_num_nodes" {
  type = number
  default = 3
}

variable "redis_node_type" {
  type = string
  default = "cache.t4g.micro"
}

variable "min_capacity" {
  type = number
  default = 1
}

variable "max_capacity" {
  type = number
  default = 3
}

