-- name: CreateCustomer :one
INSERT INTO customers (
    nik, full_name, email, phone, address, date_of_birth
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE customer_id = $1;

-- name: GetCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1;

-- name: GetCustomerAccounts :many
SELECT * FROM accounts
WHERE customer_id = $1
ORDER BY opened_date DESC;
