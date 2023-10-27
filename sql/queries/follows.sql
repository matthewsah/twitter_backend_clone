-- name: CreateFollow :one
INSERT INTO follows (id, created_at, updated_at, user_id, user_to_follow_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFollows :many
SELECT * FROM follows WHERE user_id=$1;

-- name: DeleteFollow :exec
DELETE FROM follows WHERE user_id=$1 AND user_to_follow_id=$2;