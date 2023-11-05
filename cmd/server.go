package cmd

import (
	"context"
	"os"

	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/kintuda/tech-challenge-pismo/pkg/config"
	"github.com/kintuda/tech-challenge-pismo/pkg/http"
	"github.com/kintuda/tech-challenge-pismo/pkg/middleware"
	"github.com/kintuda/tech-challenge-pismo/pkg/postgres"
	tracer "github.com/kintuda/tech-challenge-pismo/pkg/tracing"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewServerCmd(ctx context.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "Serve HTTP application",
		RunE:  StartServer,
	}

	return command
}

func StartServer(cmd *cobra.Command, arg []string) error {
	ctx := context.Background()
	app := fx.New(
		config.Module,
		postgres.Module,
		http.Module,
		middleware.Module,
		account.Module,
		transaction.Module,
		fx.Invoke(runHttpServer()),
	)

	if err := app.Start(ctx); err != nil {
		log.Error().Err(err).Msg("error starting HTTP server")
		os.Exit(1)
	}

	defer func() {
		if err := app.Stop(ctx); err != nil {
			log.Error().Err(err).Msg("error while closing app")
		}
	}()

	return nil
}

func runHttpServer() any {
	return func(
		lifecycle fx.Lifecycle, router *http.Router, cfg *config.ServerConfig, t *middleware.TransactionMiddleware) {
		lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
			if cfg.EnableTracing == "true" {
				cleanup := tracer.InitTracer(cfg)
				defer cleanup(context.Background())
			}

			router.RegisterRoutes(t)
			server := http.NewServer(router, cfg.HttpPort)
			return server.Init()
		}})
	}
}
