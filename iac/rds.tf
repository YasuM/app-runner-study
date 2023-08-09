resource "aws_db_instance" "rds" {
  instance_class         = "db.t3.micro"
  allocated_storage      = 20
  engine                 = "mysql"
  username               = "admin"
  password               = "rootroot"
  db_subnet_group_name   = aws_db_subnet_group.subnet_group.name
  vpc_security_group_ids = [aws_security_group.mysql.id]
  engine_version         = "8.0.33"
  max_allocated_storage  = "1000"
  identifier             = "database-1"
  storage_encrypted      = true
  skip_final_snapshot    = true
}

resource "aws_db_subnet_group" "subnet_group" {
  name       = local.name
  subnet_ids = [aws_subnet.private_subnet_1a.id, aws_subnet.private_subnet_1c.id]
}