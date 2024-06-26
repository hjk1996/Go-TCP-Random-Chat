provider "aws" {
  region     = var.region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}


module "iam_module" {
  source   = "./modules/iam"
  region   = var.region
  app_name = var.app_name
}


module "vpc_module" {
  source   = "./modules/vpc"
  region   = var.region
  app_name = var.app_name
  app_port = var.app_port
}

module "redis_module" {
  source                  = "./modules/redis"
  app_name                = var.app_name
  app_private_subnets     = module.vpc_module.app_private_subnets
  node_type               = var.redis_node_type
  redis_securiry_group_id = module.vpc_module.app_redis_sg_id
  num_cache_nodes         = var.redis_num_nodes
  redis_azs               = module.vpc_module.app_vpc_azs
  depends_on              = [module.vpc_module, module.iam_module]
}


module "ecs_module" {
  source                    = "./modules/ecs"
  region                    = var.region
  app_name                  = var.app_name
  app_task_sg_id            = module.vpc_module.app_task_sg_id
  app_environment_variables = var.app_environment_variables
  app_execution_role_arn    = module.iam_module.app_execution_role_arn
  app_task_role_arn         = module.iam_module.app_task_role_arn
  app_port                  = var.app_port
  // 앱의 서브넷을 프라이빗으로 하면 서비스가 정상적으로 실행안되는데, 퍼블릭 서브넷으로 바꾸니까 정상적으로 실행됨
  app_subnets               = module.vpc_module.app_public_subnets
  app_image                 = var.app_image
  redis_endpoint            = module.redis_module.redis_endpoint
  target_group_arn          = module.vpc_module.target_group_arn
  min_capacity              = var.min_capacity


  depends_on = [module.redis_module]
}

module "autoscale_module" {
  source             = "./modules/autoscale"
  autoscale_role_arn = module.iam_module.ecs_autoscale_role_arn
  min_capacity       = var.min_capacity
  max_capacity       = var.max_capacity
  service_id         = module.ecs_module.ecs_service_id
  service_name       = module.ecs_module.service_name
  cluster_name       = module.ecs_module.cluster_name
  app_name           = var.app_name

  depends_on = [module.ecs_module]

}

