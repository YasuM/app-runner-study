-- name: GetTask :one
SELECT id, name, status, created_at from task WHERE id = ? LIMIT 1;

-- name: LisTasks :many
select id, name, status, created_at from task order by created_at desc;

-- name: CreateTask :execresult
INSERT INTO task (name, status, created_at) values (?, ?, now());

-- name: UpdateTask :exec
UPDATE task SET name = ?, status = ? WHERE id = ?;

-- name: DeleteTask :exec
DELETE FROM task where id = ?;

-- name: CreateUser :execresult
INSERT INTO user(name, email, password, created_at) values(?, ?, ?, now());

-- name: CountUserByEmail :one
SELECT count(*) FROM user where email = ?;

-- name: GetUserPasswordByEmail :one
SELECT password FROM user WHERE email = ? limit 1;