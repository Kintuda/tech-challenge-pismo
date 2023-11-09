package mock

import (
	"context"

	"github.com/cockroachdb/apd"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
)

var _ transaction.TransactionRepository = (*TransactionRepositoryMock)(nil)

type TransactionRepositoryMock struct {
	Database              map[string]*transaction.Transaction
	DatabaseOperationType map[string]*transaction.OperationType
}

// ListTransactionRemainingBalance implements transaction.TransactionRepository.
func (*TransactionRepositoryMock) ListTransactionRemainingBalance(ctx context.Context, accountID string) ([]*transaction.Transaction, error) {
	panic("unimplemented")
}

// UpdateTransactionBalance implements transaction.TransactionRepository.
func (*TransactionRepositoryMock) UpdateTransactionBalance(ctx context.Context, transactionID string, accountID string, amount *apd.Decimal) error {
	panic("unimplemented")
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
