package unit

import (
	"testing"

	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
)

func TestMissingNegativeOperator(t *testing.T) {
	t.Parallel()
	service := transaction.NewTransactionService(nil, nil)

	decimal, err := service.ParseDecimal("-100", transaction.OperationType{
		OperationOperator: transaction.Positive,
	})

	if err != nil {
		t.Error(err)
	}

	if decimal.String() != "100" {
		t.Errorf("%s is not equal to 100", decimal.String())
	}
}

func TestMissingPositiveOperator(t *testing.T) {
	t.Parallel()
	service := transaction.NewTransactionService(nil, nil)

	decimal, err := service.ParseDecimal("-100", transaction.OperationType{
		OperationOperator: transaction.Negative,
	})

	if err != nil {
		t.Error(err)
	}

	if decimal.String() != "-100" {
		t.Errorf("%s is not equal to -100", decimal.String())
	}
}
