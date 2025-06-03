-- name: CreateMutation :one
INSERT INTO mutations (account_id, transfer_id, amount, previous_balance, current_balance)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListMutationsByAccount :many
SELECT
    m.id,
    m.amount,
    m.previous_balance,
    m.current_balance,
    m.created_at,
    t.reference_number,
    CASE
        WHEN t.from_account_id = m.account_id THEN 'DEBIT'
        ELSE 'CREDIT'
    END as type
FROM mutations m
JOIN transfers t ON t.id = m.transfer_id
WHERE m.account_id = (SELECT id FROM accounts WHERE account_number = $1)
ORDER BY m.created_at DESC
LIMIT $2 OFFSET $3;
