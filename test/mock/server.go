package mock

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/kintuda/tech-challenge-pismo/pkg/http"
)

type Params struct {
	TransactionController *http.TransactionController
	AccountController     *http.AccountController
}

type StaticServer struct {
	Ctx    *gin.Context
	Engine *gin.Engine
}

func CreateMockServer(w *httptest.ResponseRecorder, p Params) *StaticServer {
	c, r := gin.CreateTestContext(w)
	r.Use(http.ErrorHandler())

	if p.AccountController != nil {
		r.GET("/v1/accounts/:id", p.AccountController.GetAccount)
		r.POST("/v1/accounts", p.AccountController.CreateAccount)
	}

	if p.TransactionController != nil {
		r.POST("/v1/transactions", p.TransactionController.CreateTransaction)
	}

	return &StaticServer{
		Ctx:    c,
		Engine: r,
	}
}
