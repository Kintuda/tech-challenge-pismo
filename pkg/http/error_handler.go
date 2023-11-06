package http

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
	"github.com/rs/zerolog/log"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {

			switch {
			case errors.Is(err, io.EOF):
				c.JSON(http.StatusBadRequest, gin.H{"message": "request body is empty"})
				c.Abort()
				return

			case errors.Is(err, transaction.ErrAccountNotFound):
				validationErrors := NewHttpUnprocessableEntity()
				validationErrors.AddError(&ValidationDetails{
					Field:       "account_id",
					Value:       "",
					Constraint:  "account_not_found",
					Description: err.Error(),
				})

				c.JSON(http.StatusUnprocessableEntity, validationErrors)
				c.Abort()
				return
			case errors.Is(err, transaction.ErrTransactionTypeNotFound):
				validationErrors := NewHttpUnprocessableEntity()
				validationErrors.AddError(&ValidationDetails{
					Field:       "operation_type_id",
					Value:       "",
					Constraint:  "transaction_type_not_found",
					Description: err.Error(),
				})

				c.JSON(http.StatusUnprocessableEntity, validationErrors)
				c.Abort()
				return
			case errors.Is(err, transaction.ErrInvalidAmount):
				validationErrors := NewHttpUnprocessableEntity()
				validationErrors.AddError(&ValidationDetails{
					Field:       "amount",
					Value:       "",
					Constraint:  "invalid_amount",
					Description: err.Error(),
				})

				c.JSON(http.StatusUnprocessableEntity, validationErrors)
				c.Abort()
				return
			}

			switch errType := err.Err.(type) {
			case validator.ValidationErrors:
				validationErrors := NewHttpUnprocessableEntity()
				validationErrors.AddErrorFromField(errType)

				c.JSON(http.StatusUnprocessableEntity, validationErrors)
				c.Abort()
				return
			default:
				log.Error().Err(err.Err).Msg("internal server error")
				c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
				c.Abort()
			}
		}
	}
}
