-- name: CreateTransaction :one
INSERT INTO transactions (
    account_id,
    related_account_id,
    amount,
    transaction_type,
    description,
    reference_number
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetTransactionByID :one
SELECT * FROM transactions
WHERE transaction_id = $1;

-- name: GetAccountTransactions :many
SELECT * FROM transactions
WHERE account_id = $1
ORDER BY transaction_date DESC
LIMIT $2 OFFSET $3;

-- name: TransferFunds :exec
WITH debit AS (
    UPDATE accounts
    SET balance = balance - $1,
        updated_at = CURRENT_TIMESTAMP
    WHERE account_id = $2
    RETURNING account_id, balance
),
credit AS (
    UPDATE accounts
    SET balance = balance + $1,
        updated_at = CURRENT_TIMESTAMP
    WHERE account_id = $3
    RETURNING account_id, balance
)
INSERT INTO transactions (
    account_id,
    related_account_id,
    amount,
    transaction_type,
    description,
    reference_number,
    status
) VALUES
    ($2, $3, -$1, 'TRANSFER', $4, $5, 'COMPLETED'),
    ($3, $2, $1, 'TRANSFER', $4, $5, 'COMPLETED');
