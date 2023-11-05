package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
)

var _ transaction.TransactionRepository = (*TransactionRepositoryPg)(nil)

type TransactionRepositoryPg struct {
	PgRepository
}

func (t *TransactionRepositoryPg) GetOperationType(ctx context.Context, operationTypeID string) (*transaction.OperationType, error) {
	results := make([]*transaction.OperationType, 0)
	sql := "SELECT * FROM operation_types WHERE id = $1 AND deleted_at IS NULL LIMIT 1"

	if err := pgxscan.Select(ctx, t.GetExecutor(ctx), &results, sql, operationTypeID); err != nil {
		return nil, err
	}

	if len(results) <= 0 {
		return nil, nil
	}

	return results[0], nil
}

func (t *TransactionRepositoryPg) CreateTransaction(ctx context.Context, transaction transaction.Transaction) error {
	sql := "INSERT INTO transactions(id, account_id, operation_type_id, amount, event_date) values($1, $2, $3, $4, $5);"

	_, err := t.GetExecutor(ctx).Exec(ctx, sql, transaction.ID, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)

	return err
}
