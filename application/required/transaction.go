package required

import (
	"context"
	"fmt"
	"goplearn/ent"
)

type TransactionManager struct {
	client *ent.Client
}

func NewTransactionManager(client *ent.Client) *TransactionManager {
	return &TransactionManager{client: client}
}

func (tm *TransactionManager) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := tm.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
