output "app_execution_role_arn" {
  value = aws_iam_role.app_execution_role.arn
}

output "app_task_role_arn" {
  value = aws_iam_role.app_task_role.arn
}

output "ecs_autoscale_role_arn" {
  value = aws_iam_role.ecs_autoscale.arn
}