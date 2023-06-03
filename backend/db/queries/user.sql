-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetNumberOfUser :one
SELECT count(0)
FROM users;

-- name: CreateUser :one
INSERT INTO users ( 
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
set username = $2,
email = $3,
password = $4,
updated_at = $5
WHERE id = $1
RETURNING *;
