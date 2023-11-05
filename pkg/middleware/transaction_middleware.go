package middleware

import (
	"context"
	"net/http"

	"github.com/kintuda/tech-challenge-pismo/pkg/postgres"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type TransactionMiddleware struct {
	db *postgres.Pool
}

func NewTransactionMiddleware(db *postgres.Pool) *TransactionMiddleware {
	return &TransactionMiddleware{db: db}
}

func (t *TransactionMiddleware) OpenTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := t.db.Conn.Begin(c)

		if err != nil {
			log.Error().Err(err).Msgf("Error while starting transaction: %v", err)
			c.Status(http.StatusInternalServerError)
			c.Abort()
			return
		}

		defer func() {
			if err := recover(); err != nil || c.IsAborted() {
				log.Error().Msg("Error while processing request")

				if rbErr := tx.Rollback(c.Request.Context()); rbErr != nil {
					log.Error().Err(rbErr).Msgf("error while rollback transaction %v", err)
				}

				c.Status(http.StatusInternalServerError)
				c.Abort()
			}
		}()

		request := c.Request.WithContext(context.WithValue(c.Request.Context(), "tx", tx))

		c.Request = request

		c.Next()

		if c.IsAborted() {
			log.Error().Msg("request was aborted")

			if err := recover(); err != nil {
				log.Error().Msg("Error while recovering")
			}

			if err := tx.Rollback(c); err != nil {
				log.Error().Err(err).Msgf("Error while rolling back transaction: %v", err)
			}
			return
		}

		if err := tx.Commit(c); err != nil {
			log.Error().Err(err).Msgf("Error while commiting transaction")
			c.Abort()
		}
	}
}
