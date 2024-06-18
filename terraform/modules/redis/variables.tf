variable "app_name" {
  type = string
}

variable "redis_azs" {
  type = list(string)
}


variable "redis_securiry_group_id" {
  type = string
}

variable "num_cache_nodes" {
  type = number
}

variable "node_type" {
  type = string
}

variable "app_private_subnets" {
  type = list(string)
}


