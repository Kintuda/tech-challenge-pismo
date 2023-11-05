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
	EventDate       time.Time    `json:"event_date"`
}
