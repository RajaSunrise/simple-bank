-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount, reference_number, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: UpdateTransferStatus :one
UPDATE transfers
SET status = $2
WHERE id = $1
RETURNING *;
