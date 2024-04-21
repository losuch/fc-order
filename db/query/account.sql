-- name: GetAccountList :many
SELECT * FROM account;

-- name: GetAccountByEmail :one
SELECT * FROM account WHERE email = $1;

-- name: GetAccount :one
SELECT * FROM account WHERE account_id = $1;

-- name: CreateAccount :one
INSERT INTO account (email, hashed_password, role) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateAccount :one
UPDATE account SET hashed_password = $1, role = $2 WHERE account_id = $3 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account WHERE account_id = $1;
