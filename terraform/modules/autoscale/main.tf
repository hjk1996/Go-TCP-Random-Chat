
// autoscale의 타겟에 대한 정보
resource "aws_appautoscaling_target" "ecs_target" {
  min_capacity = var.min_capacity
  max_capacity = var.max_capacity
  resource_id        = "service/${var.cluster_name}/${var.service_name}"
  role_arn     = var.autoscale_role_arn
  // 스케일할 지표가 무엇인가?
  scalable_dimension = "ecs:service:DesiredCount"
  // 스케일링 타겟이 될 서비스의 이름
  service_namespace = "ecs"
}





// 어떤 지표를 트리거로 삼을 것인가?
resource "aws_cloudwatch_metric_alarm" "cpu_high" {
  alarm_name          = "${var.app_name}-cpu-high"
  comparison_operator = "GreaterThanOrEqualToThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ECS"
  statistic           = "Average"
  threshold           = "70"
  period              = 60


  // 어느 서비스에 대한 지표만 모티러링할 것이냐?
  dimensions = {
    ClusterName = var.cluster_name
    ServiceName = var.service_name
  }



  // 조건이 충족했을 때 어떤 행동을 할 것이냐?
  alarm_actions = [aws_appautoscaling_policy.ecs_scale_out.arn]

}


// 트리거가 발동했을 때 어떤 일을 할 것인가?
resource "aws_appautoscaling_policy" "ecs_scale_out" {
  name               = "scale-out"
  policy_type        = "StepScaling"
  resource_id        = aws_appautoscaling_target.ecs_target.resource_id
  scalable_dimension = aws_appautoscaling_target.ecs_target.scalable_dimension
  service_namespace  = aws_appautoscaling_target.ecs_target.service_namespace

  step_scaling_policy_configuration {
    adjustment_type         = "ChangeInCapacity"
    cooldown                = 60
    metric_aggregation_type = "Average"

    step_adjustment {
      scaling_adjustment          = 1
      metric_interval_lower_bound = 0
    }
  }
}




// 어떤 지표를 트리거로 삼을 것인가?
resource "aws_cloudwatch_metric_alarm" "cpu_low" {
  alarm_name          = "${var.app_name}-cpu-low"
  comparison_operator = "LessThanOrEqualToThreshold"
  evaluation_periods  = "2"
  metric_name         = "CPUUtilization"
  namespace           = "AWS/ECS"
  statistic           = "Average"
  threshold           = "30"
  period              = 60


  // 어느 서비스에 대한 지표만 모티러링할 것이냐?
  dimensions = {
    ClusterName = var.cluster_name
    ServiceName = var.service_name
  }


  // 조건이 충족했을 때 어떤 행동을 할 것이냐?
  alarm_actions = [aws_appautoscaling_policy.ecs_scale_in.arn]


}



// 트리거가 발동했을 때 어떤 일을 할 것인가?
resource "aws_appautoscaling_policy" "ecs_scale_in" {
  name               = "scale-in"
  policy_type        = "StepScaling"
  resource_id        = aws_appautoscaling_target.ecs_target.resource_id
  scalable_dimension = aws_appautoscaling_target.ecs_target.scalable_dimension
  service_namespace  = aws_appautoscaling_target.ecs_target.service_namespace

  step_scaling_policy_configuration {
    adjustment_type         = "ChangeInCapacity"
    cooldown                = 60
    metric_aggregation_type = "Average"

    step_adjustment {
      scaling_adjustment          = -1
      metric_interval_lower_bound = 0
    }
  }
}

