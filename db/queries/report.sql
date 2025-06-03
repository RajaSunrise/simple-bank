-- name: GetAccountSummary :one
SELECT
    COUNT(*) as total_transactions,
    SUM(CASE WHEN status = 'SUCCESS' THEN 1 ELSE 0 END) as success_count,
    SUM(CASE WHEN status = 'FAILED' THEN 1 ELSE 0 END) as failed_count,
    COALESCE(SUM(CASE WHEN from_account_id = $1 AND status = 'SUCCESS' THEN amount END), 0) as total_outgoing,
    COALESCE(SUM(CASE WHEN to_account_id = $1 AND status = 'SUCCESS' THEN amount END), 0) as total_incoming
FROM transfers
WHERE from_account_id = $1 OR to_account_id = $1;

-- name: DailyTransactionReport :many
SELECT
    DATE(created_at) as date,
    COUNT(*) as transaction_count,
    SUM(amount) as total_amount,
    SUM(CASE WHEN status = 'SUCCESS' THEN amount ELSE 0 END) as success_amount,
    SUM(CASE WHEN status = 'FAILED' THEN amount ELSE 0 END) as failed_amount
FROM transfers
WHERE created_at >= $1 AND created_at <= $2
GROUP BY DATE(created_at)
ORDER BY DATE(created_at) DESC;
