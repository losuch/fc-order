// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (email, hashed_password, role) VALUES ($1, $2, $3) RETURNING account_id, email, hashed_password, role, created_at
`

type CreateAccountParams struct {
	Email          string
	HashedPassword string
	Role           string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Email, arg.HashedPassword, arg.Role)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account WHERE account_id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, accountID int64) error {
	_, err := q.db.Exec(ctx, deleteAccount, accountID)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT account_id, email, hashed_password, role, created_at FROM account WHERE account_id = $1
`

func (q *Queries) GetAccount(ctx context.Context, accountID int64) (Account, error) {
	row := q.db.QueryRow(ctx, getAccount, accountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT account_id, email, hashed_password, role, created_at FROM account WHERE email = $1
`

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountList = `-- name: GetAccountList :many
SELECT account_id, email, hashed_password, role, created_at FROM account
`

func (q *Queries) GetAccountList(ctx context.Context) ([]Account, error) {
	rows, err := q.db.Query(ctx, getAccountList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.AccountID,
			&i.Email,
			&i.HashedPassword,
			&i.Role,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET hashed_password = $1, role = $2 WHERE account_id = $3 RETURNING account_id, email, hashed_password, role, created_at
`

type UpdateAccountParams struct {
	HashedPassword string
	Role           string
	AccountID      int64
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount, arg.HashedPassword, arg.Role, arg.AccountID)
	var i Account
	err := row.Scan(
		&i.AccountID,
		&i.Email,
		&i.HashedPassword,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}
