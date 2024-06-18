output "app_execution_role_arn" {
  value = aws_iam_role.app_execution_role.arn
}

output "app_task_role_arn" {
  value = aws_iam_role.app_task_role.arn
}

