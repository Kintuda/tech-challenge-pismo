package mock

import (
	"context"

	"github.com/kintuda/tech-challenge-pismo/pkg/account"
)

var _ account.AccountRepository = (*AccountRepositoryMock)(nil)

type AccountRepositoryMock struct {
	Database map[string]*account.Account
}

func NewAccountRepositoryMock() *AccountRepositoryMock {
	return &AccountRepositoryMock{
		Database: map[string]*account.Account{},
	}
}

func (a *AccountRepositoryMock) CreateAccount(ctx context.Context, account account.Account) error {
	a.Database[account.ID] = &account

	return nil
}

func (a *AccountRepositoryMock) GetAccount(ctx context.Context, ID string) (*account.Account, error) {
	content, ok := a.Database[ID]

	if ok {
		return content, nil
	}

	return nil, nil
}
