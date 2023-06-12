-- name: GetSlackChannel :one
SELECT * FROM slack_channels 
WHERE id = $1 LIMIT 1;

-- name: ListSlackChannels :many
SELECT * FROM slack_channels
ORDER BY id;

-- name: CreateSlackChannel :one
INSERT INTO slack_channels (
  id, name 
) VALUES (
  $1, $2
) 
RETURNING *;
