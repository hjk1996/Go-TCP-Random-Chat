output "redis_endpoint" {
  value = aws_elasticache_replication_group.app.primary_endpoint_address

}


