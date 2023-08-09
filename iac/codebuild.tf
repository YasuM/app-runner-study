resource "aws_codebuild_project" "db-migrate" {
  name         = "db-migrate"
  service_role = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/service-role/codebuild-a-service-role"
  source {
    type            = "GITHUB"
    location        = "https://github.com/YasuM/app-runner-study.git"
    buildspec       = "backend/buildspec.yml"
    git_clone_depth = 1
    git_submodules_config {
      fetch_submodules = false
    }
  }
  artifacts {
    type = "NO_ARTIFACTS"
  }
  vpc_config {
    security_group_ids = [
      aws_security_group.mysql.id
    ]
    subnets = [
      aws_subnet.private_subnet_1a.id,
      aws_subnet.private_subnet_1c.id,
    ]
    vpc_id = aws_vpc.vpc.id
  }
  logs_config {
    cloudwatch_logs {
      group_name = aws_cloudwatch_log_group.log.name
    }
  }

  environment {
    compute_type                = "BUILD_GENERAL1_SMALL"
    image                       = "aws/codebuild/amazonlinux2-x86_64-standard:5.0-23.07.28"
    type                        = "LINUX_CONTAINER"
    image_pull_credentials_type = "CODEBUILD"

    environment_variable {
      name  = "DB_USER"
      value = aws_db_instance.rds.username
    }
    environment_variable {
      name  = "DB_PASSWORD"
      value = aws_db_instance.rds.password
    }
    environment_variable {
      name  = "DB_HOST"
      value = aws_db_instance.rds.address
    }
  }
}

resource "aws_cloudwatch_log_group" "log" {
  name = "db-migrate"
}