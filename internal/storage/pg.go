package storage

import (
	"context"
	"database/sql"
)

type DBOps interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type TxWrapper struct {
	Tx *sql.Tx
}

func (t TxWrapper) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.Tx.QueryRowContext(ctx, query, args...).Scan(dest)
}

func (t TxWrapper) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.Tx.QueryRowContext(ctx, query, args...).Scan(dest)
}

func (t TxWrapper) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return t.Tx.ExecContext(ctx, query, args...)
}

func (t TxWrapper) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return t.Tx.QueryRowContext(ctx, query, args...)
}
