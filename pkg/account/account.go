package account

import "time"

type Account struct {
	ID             string     `json:"account_id"`
	DocumentNumber string     `json:"document_number"`
	CreatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
}
