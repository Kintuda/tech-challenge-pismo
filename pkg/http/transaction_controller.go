package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
	"github.com/rs/zerolog/log"
)

type TransactionController struct {
	service *transaction.TransactionService
}

type CreateTransaction struct {
	AccountID       string `json:"account_id" binding:"required"`
	OperationTypeID string `json:"operation_type_id" binding:"required"`
	Amount          string `json:"amount" binding:"required"`
}

func NewTransactionController(service *transaction.TransactionService) *TransactionController {
	return &TransactionController{
		service: service,
	}
}

func (t *TransactionController) CreateTransaction(c *gin.Context) {
	payload := new(CreateTransaction)

	if err := c.ShouldBindBodyWith(payload, binding.JSON); err != nil {
		if err := c.Error(err); err != nil {
			log.Error().Err(err).Msgf("error while handling error")
		}

		c.Abort()
		return
	}

	account, err := t.service.CreateTransaction(c.Request.Context(), payload.OperationTypeID, payload.Amount, payload.AccountID)

	if err != nil {
		if err := c.Error(err); err != nil {
			log.Err(err).Msgf("error while handling error %v", err)
		}
		c.Abort()
		return
	}

	c.JSON(201, account)
}
