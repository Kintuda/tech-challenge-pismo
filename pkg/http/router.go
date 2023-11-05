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
	engine      *gin.Engine
	account     *AccountController
	transaction *TransactionController
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

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	validators := map[string]validator.Func{
	// 		"brazilianDocumentNumber": helper.BrazilianDocumentNumber,
	// 	}

	// 	for key, value := range validators {
	// 		if err := v.RegisterValidation(key, value); err != nil {
	// 			log.Error().Msgf("cannot register validation %s", key)
	// 		}
	// 	}

	// 	v.RegisterCustomTypeFunc(helper.ValidateJSONDateType, carbon.Date{})
	// }

	return &Router{
		engine:      router,
		account:     accountController,
		transaction: transactionController,
	}
}

func (r *Router) RegisterRoutes(t *middleware.TransactionMiddleware) {
	// r.engine.Use(ErrorHandler())
	v1 := r.engine.Group("/v1")

	v1.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{})
	})

	account := v1.Group("/accounts")
	account.GET("/:id", r.account.GetAccount)
	account.POST("/", t.OpenTransaction(), r.account.CreateAccount)

	transaction := v1.Group("/transactions")
	transaction.POST("/", r.transaction.CreateTransaction)
}
