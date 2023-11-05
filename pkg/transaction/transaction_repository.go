package transaction

import "context"

type TransactionRepository interface {
	GetOperationType(ctx context.Context, operationTypeID string) (*OperationType, error)
	CreateTransaction(ctx context.Context, transaction Transaction) error
}
