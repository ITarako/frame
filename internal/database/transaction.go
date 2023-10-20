package database

import (
	"context"
	"database/sql"
	"errors"
)

type TransactionManager interface {
	Do(context.Context, func(context.Context, *sql.Tx) error) error
}

type transactionManagerSql struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *transactionManagerSql {
	return &transactionManagerSql{db: db}
}

func (tm *transactionManagerSql) Do(ctx context.Context, fn func(context.Context, *sql.Tx) error) (err error) {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		errRollback := tx.Rollback()
		if errRollback != nil && !errors.Is(errRollback, sql.ErrTxDone) {
			err = errors.Join(err, errRollback)
		}
	}()

	if err = fn(ctx, tx); err != nil {
		return
	}

	return tx.Commit()
}
