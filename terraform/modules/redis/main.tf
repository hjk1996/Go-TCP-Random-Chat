resource "aws_elasticache_subnet_group" "main" {
  name       = "${var.app_name}-subnet-group"
  subnet_ids = var.app_private_subnets

}

resource "aws_elasticache_replication_group" "app" {
  automatic_failover_enabled  = var.num_cache_nodes >= 2 ? true : false
  preferred_cache_cluster_azs = slice(var.redis_azs, 0, var.num_cache_nodes)
  replication_group_id        = "${var.app_name}-rep-group-1"
  description                 = "hello"
  node_type                   = var.node_type
  num_cache_clusters          = var.num_cache_nodes
  subnet_group_name           = aws_elasticache_subnet_group.main.name
  port                        = 6379

  lifecycle {
    ignore_changes = [num_cache_clusters]
  }
}



# resource "aws_elasticache_cluster" "replica" {
#   count = 1
#   cluster_id           = "${var.app_name}-rep-group-1-${count.index}"
#   replication_group_id = aws_elasticache_replication_group.app.id
# }