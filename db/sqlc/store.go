package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	UpdateAccountTx(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

// SQLStore provides all functions to execute db queries and transaction
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}




// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// UpdateUserTx updates user within single transaction
func (store *SQLStore) UpdateAccountTx(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	var result Account

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result, err = q.UpdateAccount(ctx, arg)
		return err
	})
	return result, err
}
