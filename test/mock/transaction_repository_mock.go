package mock

import (
	"context"

	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
)

var _ transaction.TransactionRepository = (*TransactionRepositoryMock)(nil)

type TransactionRepositoryMock struct {
	Database              map[string]*transaction.Transaction
	DatabaseOperationType map[string]*transaction.OperationType
}

func NewTransactionRepositoryMock() *TransactionRepositoryMock {
	return &TransactionRepositoryMock{
		Database:              map[string]*transaction.Transaction{},
		DatabaseOperationType: map[string]*transaction.OperationType{},
	}
}

func (t *TransactionRepositoryMock) CreateTransaction(ctx context.Context, transaction transaction.Transaction) error {
	t.Database[transaction.ID] = &transaction

	return nil
}

func (t *TransactionRepositoryMock) GetOperationType(ctx context.Context, operationTypeID string) (*transaction.OperationType, error) {
	content, ok := t.DatabaseOperationType[operationTypeID]

	if ok {
		return content, nil
	}

	return nil, nil
}
