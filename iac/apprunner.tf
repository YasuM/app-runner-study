resource "aws_apprunner_service" "apprunner" {
  service_name = local.name
  source_configuration {
    authentication_configuration {
      access_role_arn = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/service-role/AppRunnerECRAccessRole"
    }
    image_repository {
      image_configuration {
        port = "8080"
        runtime_environment_variables = {
          MYSQL_USER     = "admin"
          MYSQL_PASSWORD = "rootroot"
          MYSQL_HOST     = aws_db_instance.rds.address
        }
      }
      image_identifier      = "${aws_ecr_repository.ecr.repository_url}:v10"
      image_repository_type = "ECR"
    }
    auto_deployments_enabled = false
  }
  network_configuration {
    egress_configuration {
      egress_type       = "VPC"
      vpc_connector_arn = aws_apprunner_vpc_connector.connector.arn
    }
  }
  health_check_configuration {
    protocol = "HTTP"
    path     = "/health"
  }
}

resource "aws_apprunner_vpc_connector" "connector" {
  vpc_connector_name = local.name
  subnets            = [aws_subnet.private_subnet_1a.id, aws_subnet.private_subnet_1c.id]
  security_groups    = [aws_security_group.mysql.id]
}