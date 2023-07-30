data "aws_caller_identity" "current" {}

locals {
  name = "app-runner-study"
}
data "aws_db_instance" "database" {
  db_instance_identifier = "database-1"
}

resource "aws_ecr_repository" "ecr" {
  name                 = local.name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = false
  }
}

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
      image_identifier      = "${aws_ecr_repository.ecr.repository_url}:v9"
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

resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true

}

resource "aws_subnet" "private_subnet_1a" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "10.0.3.0/24"
  availability_zone = "ap-northeast-1a"
  tags = {
    Name = "subnet-private-1a"
  }
}

resource "aws_subnet" "private_subnet_1c" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "10.0.4.0/24"
  availability_zone = "ap-northeast-1c"
  tags = {
    Name = "subnet-private-1c"
  }
}

resource "aws_security_group" "mysql" {
  name   = "allow_mysql"
  vpc_id = aws_vpc.vpc.id

  ingress {
    description = "MySQL from VPC"
    from_port   = 3306
    to_port     = 3306
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_db_instance" "rds" {
  instance_class = "db.t3.micro"
  allocated_storage = 20
  engine = "mysql"
  db_subnet_group_name = aws_db_subnet_group.subnet_group.name
  engine_version = "8.0.33"
  max_allocated_storage = "1000"
  identifier = "database-1"
  storage_encrypted = true
  skip_final_snapshot = true
}

resource "aws_db_subnet_group" "subnet_group" {
  name       = local.name
  subnet_ids = [aws_subnet.private_subnet_1a.id, aws_subnet.private_subnet_1c.id]
}