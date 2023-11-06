package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	httpServer "github.com/kintuda/tech-challenge-pismo/pkg/http"

	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/kintuda/tech-challenge-pismo/test/mock"
)

func TestInvalidID(t *testing.T) {
	t.Parallel()

	accountRepo := mock.NewAccountRepositoryMock()
	accountService := account.NewAccountService(accountRepo)

	accountController := httpServer.NewAccountController(accountService)

	w := httptest.NewRecorder()

	mockServer := mock.CreateMockServer(w, mock.Params{
		AccountController: accountController,
	})

	mockServer.Ctx.Request, _ = http.NewRequest(http.MethodGet, "/v1/accounts/invalid_uuid", nil)

	mockServer.Engine.ServeHTTP(w, mockServer.Ctx.Request)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestMissingFields(t *testing.T) {
	t.Parallel()

	accountRepo := mock.NewAccountRepositoryMock()
	accountService := account.NewAccountService(accountRepo)

	accountController := httpServer.NewAccountController(accountService)

	w := httptest.NewRecorder()

	mockServer := mock.CreateMockServer(w, mock.Params{
		AccountController: accountController,
	})

	mockServer.Ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/accounts", bytes.NewReader([]byte("{}")))

	mockServer.Engine.ServeHTTP(w, mockServer.Ctx.Request)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	errorResponse := &httpServer.HttpUnprocessableEntity{}

	if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
		t.Error(err)
	}

	if errorResponse.Errors[0].Field != "document_number" {
		t.Errorf("missing validation from required field document_number")
	}
}
