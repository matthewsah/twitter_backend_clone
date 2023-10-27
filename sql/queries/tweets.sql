-- name: CreateTweet :one
INSERT INTO tweets (id, created_at, updated_at, content, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
