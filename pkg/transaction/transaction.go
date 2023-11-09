package transaction

import (
	"time"

	"github.com/cockroachdb/apd"
)

type Transaction struct {
	ID              string       `json:"id"`
	AccountID       string       `json:"account_id"`
	OperationTypeID string       `json:"operation_type_id"`
	Amount          *apd.Decimal `json:"amount"`
	Balance         *apd.Decimal `json:"balance"`
	EventDate       time.Time    `json:"event_date"`
}

func (t *Transaction) SubRemainingBalance(amount *apd.Decimal) error {
	value := new(apd.Decimal)
	ctx := apd.BaseContext.WithPrecision(5)

	// 60 - 23.5
	_, err := ctx.Sub(value, amount, t.Balance)

	if err != nil {
		return err
	}

	t.Balance = value

	return nil
}
