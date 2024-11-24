package db

import (
	"fmt"
)

type TransactionManager interface {
	RunInTransaction(fn func(TransactionalDbClient) error) error
}

type txManager struct {
	db TransactionalDbClient
}

func NewTransactionManager(db TransactionalDbClient) TransactionManager {
	return &txManager{db: db}
}

func (m *txManager) RunInTransaction(fn func(TransactionalDbClient) error) error {
	tx, err := m.db.BeginTx()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
