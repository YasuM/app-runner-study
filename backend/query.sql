-- name: GetTask :one
SELECT id, name, status, created_at from task WHERE id = ? LIMIT 1;

-- name: LisTasks :many
select id, name, status, created_at from task order by created_at desc;

-- name: CreateTask :execresult
INSERT INTO task (name, status, created_at) values (?, ?, now());

-- name: UpdateTask :exec
UPDATE task SET name = ?, status = ? WHERE id = ?;