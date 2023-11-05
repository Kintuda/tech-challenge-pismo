package account

import "context"

type AccountRepository interface {
	GetAccount(ctx context.Context, ID string) (*Account, error)
	CreateAccount(ctx context.Context, account Account) error
}
