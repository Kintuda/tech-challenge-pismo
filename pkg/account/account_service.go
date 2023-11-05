package account

import (
	"context"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) GetAccount(ctx context.Context, ID string) (*Account, error) {
	return a.repo.GetAccount(ctx, ID)
}

func (a *AccountService) OnlyNumbers(data string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(data, "")
}

func (a *AccountService) CreateAccount(ctx context.Context, documentNumber string) (*Account, error) {
	cleanDocumentNumber := a.OnlyNumbers(documentNumber)

	account := Account{
		ID:             uuid.NewString(),
		CreatedAt:      time.Now(),
		DocumentNumber: cleanDocumentNumber,
		DeletedAt:      nil,
	}

	if err := a.repo.CreateAccount(ctx, account); err != nil {
		log.Error().Err(err).Msg("error while creating account")
		return nil, err
	}

	return &account, nil
}
