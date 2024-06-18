
output "app_private_subnets" {
  value = module.vpc.private_subnets
}

output "app_vpc_azs" {
  value = module.vpc.azs
}

output "app_vpc_id" {
  value = module.vpc.vpc_id
}

output "app_task_sg_id" {
  value = aws_security_group.app_task_sg.id
}

output "app_redis_sg_id" {
  value = aws_security_group.app_redis_sg.id
}

