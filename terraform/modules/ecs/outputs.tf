output "cluster_name" {
  value = aws_ecs_cluster.main.name
}


output "service_name" {
  value = aws_ecs_service.app_ecs_service.name
}

output "ecs_service_id" {
  value = aws_ecs_service.app_ecs_service.id
}

