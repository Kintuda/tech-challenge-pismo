package middleware

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("x-request-id")

		if requestID == "" {
			requestID = uuid.NewString()
			c.Request.Header.Add("x-request-id", requestID)
		}

		request := c.Request.WithContext(context.WithValue(c.Request.Context(), "x-request-id", requestID))
		c.Request = request

		c.Header("x-request-id", requestID)
		c.Next()
	}
}
