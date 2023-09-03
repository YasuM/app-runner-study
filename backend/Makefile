up: 
	docker-compose up -d
	sql-migrate up
down:
	docker-compose down
db-migrate:
	sql-migrate up
sqlc-generate: db-migrate
	mysqldump -h 127.0.0.1 -u root -proot -d task --ignore-table=task.migrations > schema.sql
	sqlc generate
