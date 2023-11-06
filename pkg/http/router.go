package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kintuda/tech-challenge-pismo/pkg/config"
	"github.com/kintuda/tech-challenge-pismo/pkg/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const (
	production = "production"
)

type Router struct {
	Engine      *gin.Engine
	Account     *AccountController
	Transaction *TransactionController
}

func NewRouter(
	cfg *config.ServerConfig,
	accountController *AccountController,
	transactionController *TransactionController,
) *Router {
	if cfg.Env == production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(
		gin.Recovery(),
		middleware.RequestIDMiddleware(),
	)

	if cfg.EnableTracing == "true" {
		router.Use(otelgin.Middleware("tech-challenge-pismo-server"))
	}

	return &Router{
		Engine:      router,
		Account:     accountController,
		Transaction: transactionController,
	}
}

func (r *Router) RegisterRoutes(t *middleware.TransactionMiddleware) {
	r.Engine.Use(ErrorHandler())
	v1 := r.Engine.Group("/v1")

	v1.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})

	account := v1.Group("/accounts")
	account.GET("/:id", r.Account.GetAccount)
	account.POST("/", t.OpenTransaction(), r.Account.CreateAccount)

	transaction := v1.Group("/transactions")
	transaction.POST("/", r.Transaction.CreateTransaction)
}
