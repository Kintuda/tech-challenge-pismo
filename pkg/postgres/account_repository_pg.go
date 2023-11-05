package postgres

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/kintuda/tech-challenge-pismo/pkg/account"
)

var _ account.AccountRepository = (*AccountRepositoryPg)(nil)

type AccountRepositoryPg struct {
	PgRepository
}

func (a *AccountRepositoryPg) CreateAccount(ctx context.Context, account account.Account) error {
	sql := `INSERT INTO accounts(
		id,
		document_number,
		created_at,
		deleted_at
	) VALUES($1, $2, $3, $4)`

	_, err := a.GetExecutor(ctx).Exec(ctx, sql,
		account.ID,
		account.DocumentNumber,
		account.CreatedAt,
		account.DeletedAt,
	)

	return err
}

func (a *AccountRepositoryPg) GetAccount(ctx context.Context, ID string) (*account.Account, error) {
	results := make([]*account.Account, 0)
	query := "SELECT * FROM accounts WHERE id = $1 AND deleted_at IS NULL"

	if err := pgxscan.Select(ctx, a.GetExecutor(ctx), &results, query, ID); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return results[0], nil
}
