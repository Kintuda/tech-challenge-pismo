package transaction

import (
	"context"
)

type TransactionRepository interface {
	GetOperationType(ctx context.Context, operationTypeID string) (*OperationType, error)
	CreateTransaction(ctx context.Context, transaction Transaction) error
	UpdateTransactionBalance(ctx context.Context, transactionID, accountID string, amount string) error
	ListTransactionRemainingBalance(ctx context.Context, accountID string) ([]*Transaction, error)
}
