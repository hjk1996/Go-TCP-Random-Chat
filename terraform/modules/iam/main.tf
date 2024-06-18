



resource "aws_iam_role" "app_task_role" {
  name = "${var.app_name}-task-role"
  // 신뢰 정책 설정 (어떤 서비스가 iam role을 맡을 수 있는가?)
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Effect = "Allow",
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      },
      Action = "sts:AssumeRole"
    }]
  })
}


data "aws_caller_identity" "current" {

}

resource "aws_iam_policy" "app_task_role_policy" {
  name = "${var.app_name}-role-policy"
  policy = jsonencode(
    {
      Version = "2012-10-17",
      Statement = [
        {
          Effect = "Allow",
          Action = [
            "ssm:GetParameter",
            "ssm:GetParameters",
            "ssm:GetParameterHistory",
            "ssm:GetParametersByPath"
          ],
          Resource = [
            "arn:aws:ssm:${var.region}:${data.aws_caller_identity.current.account_id}:parameter/*",
          ]
        }
      ]

    },


  )

}

resource "aws_iam_role_policy_attachment" "task_role_attachment_1" {
  role       = aws_iam_role.app_task_role.name
  policy_arn = aws_iam_policy.app_task_role_policy.arn
}



/////////////////////////



resource "aws_iam_role" "app_execution_role" {
  name = "${var.app_name}-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        },
        Action = "sts:AssumeRole"
      }
    ]
  })
}

// 기본적인 ECS Task Execution Role
resource "aws_iam_role_policy_attachment" "execution_role_policy_1" {
  role       = aws_iam_role.app_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

// 클라우드 워치에 로그 쓸 수 있는 정책
resource "aws_iam_role_policy_attachment" "execution_role_policy_2" {
  role       = aws_iam_role.app_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}