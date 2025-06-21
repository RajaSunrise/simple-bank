-- name: CreateAccount :one
INSERT INTO accounts (
    customer_id, branch_code, balance, account_number, account_type, opened_date
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetAccountByNumber :one
SELECT * FROM accounts
WHERE account_number = $1;

-- name: UpdateAccountBalance :exec
UPDATE accounts
SET
    balance = balance + $1,
    updated_at = CURRENT_TIMESTAMP
WHERE account_id = $2;
