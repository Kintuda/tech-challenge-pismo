package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	httpServer "github.com/kintuda/tech-challenge-pismo/pkg/http"
	"github.com/kintuda/tech-challenge-pismo/pkg/transaction"

	"github.com/kintuda/tech-challenge-pismo/pkg/account"
	"github.com/kintuda/tech-challenge-pismo/test/mock"
)

func TestMissingAmountField(t *testing.T) {
	t.Parallel()

	accountRepo := mock.NewAccountRepositoryMock()
	transactionRepo := mock.NewTransactionRepositoryMock()
	accountID := uuid.NewString()
	operationID := uuid.NewString()

	transactionRepo.DatabaseOperationType[operationID] = &transaction.OperationType{
		ID:                operationID,
		Description:       "testing",
		OperationOperator: transaction.Positive,
		CreatedAt:         time.Now(),
		DeletedAt:         nil,
	}

	accountService := account.NewAccountService(accountRepo)

	transactionService := transaction.NewTransactionService(transactionRepo, accountService)
	transactionController := httpServer.NewTransactionController(transactionService)

	w := httptest.NewRecorder()

	mockServer := mock.CreateMockServer(w, mock.Params{
		TransactionController: transactionController,
	})

	err := accountRepo.CreateAccount(context.TODO(), account.Account{
		ID:             accountID,
		DocumentNumber: "70940932091", //fake generated CPF,
		CreatedAt:      time.Now(),
		DeletedAt:      nil,
	})

	if err != nil {
		t.Error(err)
	}

	jsonBody, err := json.Marshal(httpServer.CreateTransaction{
		AccountID:       accountID,
		OperationTypeID: operationID,
	})

	if err != nil {
		t.Error(err)
	}

	mockServer.Ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewReader(jsonBody))

	mockServer.Engine.ServeHTTP(w, mockServer.Ctx.Request)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	errorResponse := &httpServer.HttpUnprocessableEntity{}

	if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
		t.Error(err)
	}

	if errorResponse.Errors[0].Field != "amount" {
		t.Errorf("missing validation from required field amount")
	}
}

func TestAccountNotFound(t *testing.T) {
	t.Parallel()

	accountRepo := mock.NewAccountRepositoryMock()
	transactionRepo := mock.NewTransactionRepositoryMock()
	operationID := uuid.NewString()

	transactionRepo.DatabaseOperationType[operationID] = &transaction.OperationType{
		ID:                operationID,
		Description:       "testing",
		OperationOperator: transaction.Positive,
		CreatedAt:         time.Now(),
		DeletedAt:         nil,
	}

	accountService := account.NewAccountService(accountRepo)

	transactionService := transaction.NewTransactionService(transactionRepo, accountService)
	transactionController := httpServer.NewTransactionController(transactionService)

	w := httptest.NewRecorder()

	mockServer := mock.CreateMockServer(w, mock.Params{
		TransactionController: transactionController,
	})

	jsonBody, err := json.Marshal(httpServer.CreateTransaction{
		AccountID:       uuid.NewString(),
		OperationTypeID: operationID,
		Amount:          "0",
	})

	if err != nil {
		t.Error(err)
	}

	mockServer.Ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewReader(jsonBody))

	mockServer.Engine.ServeHTTP(w, mockServer.Ctx.Request)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	errorResponse := &httpServer.HttpUnprocessableEntity{}

	if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
		t.Error(err)
	}

	if errorResponse.Errors[0].Description != "account not found" {
		t.Errorf("missing validation from account not found")
	}
}

func TestMissingOperationType(t *testing.T) {
	t.Parallel()

	accountRepo := mock.NewAccountRepositoryMock()
	transactionRepo := mock.NewTransactionRepositoryMock()
	accountID := uuid.NewString()

	accountService := account.NewAccountService(accountRepo)

	transactionService := transaction.NewTransactionService(transactionRepo, accountService)
	transactionController := httpServer.NewTransactionController(transactionService)

	w := httptest.NewRecorder()

	mockServer := mock.CreateMockServer(w, mock.Params{
		TransactionController: transactionController,
	})

	err := accountRepo.CreateAccount(context.TODO(), account.Account{
		ID:             accountID,
		DocumentNumber: "70940932091", //fake generated CPF,
		CreatedAt:      time.Now(),
		DeletedAt:      nil,
	})

	if err != nil {
		t.Error(err)
	}

	jsonBody, err := json.Marshal(httpServer.CreateTransaction{
		AccountID:       accountID,
		OperationTypeID: uuid.NewString(),
		Amount:          "0",
	})

	if err != nil {
		t.Error(err)
	}

	mockServer.Ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/transactions", bytes.NewReader(jsonBody))

	mockServer.Engine.ServeHTTP(w, mockServer.Ctx.Request)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	errorResponse := &httpServer.HttpUnprocessableEntity{}

	if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
		t.Error(err)
	}

	if errorResponse.Errors[0].Description != "transaction type not found" {
		t.Errorf("missing validation from transaction type not found")
	}
}
