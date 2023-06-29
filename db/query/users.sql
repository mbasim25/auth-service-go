-- name: CreateUser :one
INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, created_at FROM users ORDER BY id LIMIT $1 OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
