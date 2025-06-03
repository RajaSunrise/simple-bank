-- name: CreateAccount :one
INSERT INTO accounts (full_name, account_number, balance)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByNumber :one
SELECT * FROM accounts
WHERE account_number = $1 LIMIT 1;

-- name: UpdateBalance :one
UPDATE accounts
SET balance = balance + $2
WHERE id = $1
RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: DeleteAccount :exec
DELETE FROM accounts 
WHERE id = $1;
