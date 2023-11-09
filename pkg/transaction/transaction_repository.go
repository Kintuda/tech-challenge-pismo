package transaction

import (
	"context"

	"github.com/cockroachdb/apd"
)

type TransactionRepository interface {
	GetOperationType(ctx context.Context, operationTypeID string) (*OperationType, error)
	CreateTransaction(ctx context.Context, transaction Transaction) error
	UpdateTransactionBalance(ctx context.Context, transactionID, accountID string, amount *apd.Decimal) error
	ListTransactionRemainingBalance(ctx context.Context, accountID string) ([]*Transaction, error)
}
