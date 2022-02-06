-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    user_id, name, email, roles, password_hash, date_created, date_updated
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

