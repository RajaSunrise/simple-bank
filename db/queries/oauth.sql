

-- name: CreateUser :one
INSERT INTO users (username, password_hash, email, account_id, role)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByAccount :one
SELECT * FROM users
WHERE account_id = $1 LIMIT 1;
