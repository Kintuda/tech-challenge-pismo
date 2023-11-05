package transaction

import "time"

type Operator string

const (
	Positive Operator = "POSITIVE"
	Negative Operator = "NEGATIVE"
)

type OperationType struct {
	ID                string     `json:"operation_type_id"`
	Description       string     `json:"description"`
	OperationOperator Operator   `json:"operation_operator"`
	CreatedAt         time.Time  `json:"created_at"`
	DeletedAt         *time.Time `json:"-"`
}
