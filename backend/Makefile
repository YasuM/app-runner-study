app-up:
	docker-compose up -d
db-migrate:
	sql-migrate up
sqlc-generate: db-migrate
	mysqldump -h 127.0.0.1 -u root -proot -d task --ignore-table=task.migrations > schema.sql
	sqlc generate