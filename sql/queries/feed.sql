-- name: GetFeed :many
SELECT users.username, tweets.created_at, tweets.content
FROM follows
LEFT JOIN tweets ON follows.user_to_follow_id = tweets.user_id
LEFT JOIN users ON follows.user_to_follow_id = users.id
WHERE follows.user_id = $1
ORDER BY tweets.created_at DESC;