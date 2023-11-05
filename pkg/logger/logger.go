package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	Instance zerolog.Logger
}

func NewLoggerWithContext(ctx context.Context) Logger {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		contextValue, ok := ctx.Value("x-request-id").(string)

		if ok {
			return c.Str("x-request-id", contextValue)
		}

		return c
	})

	return Logger{Instance: logger}
}
