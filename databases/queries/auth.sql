-- name: RegisterUser :one
INSERT INTO auth_users (
    customer_id, username, password_hash, role
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM auth_users
WHERE username = $1;
