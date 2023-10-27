-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name, username, password)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;