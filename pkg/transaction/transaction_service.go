package transaction

import (
	"context"
	"errors"
	"time"

	"github.com/cockroachdb/apd"
	"github.com/google/uuid"
	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/kintuda/tech-challenge-pismo/pkg/logger"
	"github.com/rs/zerolog/log"
)

var (
	ErrAccountNotFound         = errors.New("account not found")
	ErrInvalidAmount           = errors.New("invalid amount")
	ErrTransactionTypeNotFound = errors.New("transaction type not found")
)

type TransactionService struct {
	accountService *account.AccountService
	repo           TransactionRepository
}

func NewTransactionService(repo TransactionRepository, accountService *account.AccountService) *TransactionService {
	return &TransactionService{
		repo:           repo,
		accountService: accountService,
	}
}

func (t *TransactionService) ParseDecimal(amount string, operationType OperationType) (*apd.Decimal, error) {
	parsedAmount := new(apd.Decimal)

	if amount == "" {
		_, err := parsedAmount.SetFloat64(0)

		if err != nil {
			return nil, err
		}

		return parsedAmount, nil
	}

	if _, _, err := parsedAmount.SetString(amount); err != nil {
		log.Error().Err(err).Msgf("%s is a invalid amount", amount)
		return nil, err
	}

	if (parsedAmount.Negative && operationType.OperationOperator == Negative) || (!parsedAmount.Negative && operationType.OperationOperator == Positive) {
		return parsedAmount, nil
	} else {
		return parsedAmount.Neg(parsedAmount), nil
	}
}

func (t *TransactionService) CreateTransaction(ctx context.Context, operationTypeID string, amount string, accountID string) (*Transaction, error) {
	log := logger.NewLoggerWithContext(ctx)
	account, err := t.accountService.GetAccount(ctx, accountID)

	if err != nil {
		log.Instance.Error().Err(err).Msg("error while getting account")
		return nil, err
	}

	if account == nil {
		log.Instance.Warn().Msg("account not found")
		return nil, ErrAccountNotFound
	}

	operationType, err := t.repo.GetOperationType(ctx, operationTypeID)

	if err != nil {
		log.Instance.Error().Err(err).Msg("error while getting transaction type")
		return nil, err
	}

	if operationType == nil {
		log.Instance.Warn().Msg("transaction type not found")
		return nil, ErrTransactionTypeNotFound
	}

	parsedAmount, err := t.ParseDecimal(amount, *operationType)

	if err != nil {
		log.Instance.Error().Err(err).Msg("invalid amount")
		return nil, ErrInvalidAmount
	}

	transaction := Transaction{
		ID:              uuid.NewString(),
		AccountID:       account.ID,
		OperationTypeID: operationType.ID,
		Amount:          parsedAmount,
		EventDate:       time.Now(),
		Balance:         parsedAmount,
	}

	if operationType.OperationOperator == Positive {
		pendingTransactions, err := t.repo.ListTransactionRemainingBalance(ctx, accountID)

		if err != nil {
			log.Instance.Error().Err(err).Msg("error while searching pending transaction")
			return nil, err
		}

		if len(pendingTransactions) <= 0 {
			if err := t.repo.CreateTransaction(ctx, transaction); err != nil {
				log.Instance.Error().Err(err).Msg("error while processing transaction")
				return nil, err
			}
		}

		remainingBalance := new(apd.Decimal)
		remainingBalance.Set(transaction.Amount)

		for _, trx := range pendingTransactions {
			if trx.Balance.Negative {
				if remainingBalance.Cmp(trx.Balance) > 0 {
					if err := t.repo.UpdateTransactionBalance(ctx, trx.ID, trx.AccountID, "0"); err != nil {
						log.Instance.Error().Err(err).Msg("error while processing transaction")
						return nil, err
					}

					dst := new(apd.Decimal)
					newBalance := apd.BaseContext.WithPrecision(5)

					if _, err = newBalance.Sub(dst, remainingBalance, trx.Balance); err != nil {
						log.Instance.Error().Err(err).Msg("error while sub new balance")
						return nil, err
					}

					remainingBalance.Set(dst)
				} else {
					//partial settlement
				}
			}

			//Insert transaction and update in the database the old transaction balance, and update the remaining balance
			transaction.Balance = remainingBalance

			if err := t.repo.CreateTransaction(ctx, transaction); err != nil {
				log.Instance.Error().Err(err).Msg("error while processing transaction")
				return nil, err
			}

			return &transaction, nil
		}
	} else {
		if err := t.repo.CreateTransaction(ctx, transaction); err != nil {
			log.Instance.Error().Err(err).Msg("error while processing transaction")
			return nil, err
		}
	}

	return &transaction, nil
}
