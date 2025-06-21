-- name: RecordAuditLog :one
INSERT INTO audit_logs (
    user_id, action, description, ip_address, user_agent
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;
