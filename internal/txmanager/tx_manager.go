package txmanager

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"test/internal/storage"
)

const txkey = "manager.tx.key"

type TxManager struct {
	db *sqlx.DB
}

func NewTxManager(db *sqlx.DB) *TxManager {
	return &TxManager{db: db}
}

func (m *TxManager) injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txkey, tx)
}

func (m *TxManager) extractTx(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(txkey).(*sql.Tx)
	return tx, ok
}

func (m *TxManager) GetTxOrDefault(ctx context.Context) storage.DBOps {
	if tx, ok := m.extractTx(ctx); ok {
		return storage.TxWrapper{Tx: tx}
	}
	return m.db

}

func (m *TxManager) Do(ctx context.Context, f func(ctx context.Context) error) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Commit()
	}()
	if err := f(m.injectTx(ctx, tx)); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
