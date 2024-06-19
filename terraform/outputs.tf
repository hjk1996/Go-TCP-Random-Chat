

output "vpc_id" {
  value = module.vpc_module.app_vpc_id
}

output "private_subnet_ids" {
  value = module.vpc_module.app_private_subnets
}

output "redis_endpoint" {
  value = module.redis_module.redis_endpoint
}

output "lb_dns_name" {
  value = module.vpc_module.lb_dns_name
}


