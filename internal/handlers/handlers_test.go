package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/pkg"
)

// MockAccountRepository
type MockAccountRepository struct {
	CreateAccountFunc     func(ctx context.Context, accountId *int64, initialBalance *string) error
	GetAccountBalanceFunc func(ctx context.Context, accountId int64) (string, string, error)
}

func (m *MockAccountRepository) CreateAccount(ctx context.Context, accountId *int64, initialBalance *string) error {
	return m.CreateAccountFunc(ctx, accountId, initialBalance)
}

func (m *MockAccountRepository) GetAccountBalance(ctx context.Context, accountId int64) (string, string, error) {
	return m.GetAccountBalanceFunc(ctx, accountId)
}

// MockTransactionsRepository
type MockTransactionsRepository struct {
	CreateTransactionFunc func(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error
}

func (m *MockTransactionsRepository) CreateTransaction(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {
	return m.CreateTransactionFunc(ctx, sourceAccountId, destinationAccountId, amount)
}

func TestCreateAccountHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockAccountRepository{
			CreateAccountFunc: func(ctx context.Context, accountId *int64, initialBalance *string) error {
				return nil
			},
		}

		handler := &accountHandler{accountRespository: mockRepo}
		router := gin.Default()
		router.POST("/accounts", handler.CreateAccountHandler)

		body := pkg.Account{
			AccountID:      123,
			InitialBalance: "100.00",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != 201 {
			t.Errorf("Expected status 201, got %d", w.Code)
		}
	})

	t.Run("Invalid Input", func(t *testing.T) {
		mockRepo := &MockAccountRepository{}
		handler := &accountHandler{accountRespository: mockRepo}
		router := gin.Default()
		router.POST("/accounts", handler.CreateAccountHandler)

		body := pkg.Account{
			AccountID:      0, // Invalid
			InitialBalance: "100.00",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != 400 {
			t.Errorf("Expected status 400, got %d", w.Code)
		}
	})
}

func TestCreateTransactionHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockRepo := &MockTransactionsRepository{
			CreateTransactionFunc: func(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {
				return nil
			},
		}

		handler := &transactionsHandler{transactionRespository: mockRepo}
		router := gin.Default()
		router.POST("/transactions", handler.CreateTransactionHandler)

		body := pkg.TransactionsService{
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               "50.00",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != 200 {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	t.Run("Insufficient Funds", func(t *testing.T) {
		mockRepo := &MockTransactionsRepository{
			CreateTransactionFunc: func(ctx context.Context, sourceAccountId *int64, destinationAccountId *int64, amount *string) error {
				return errors.New("insufficient funds")
			},
		}

		handler := &transactionsHandler{transactionRespository: mockRepo}
		router := gin.Default()
		router.POST("/transactions", handler.CreateTransactionHandler)

		body := pkg.TransactionsService{
			SourceAccountID:      1,
			DestinationAccountID: 2,
			Amount:               "5000.00",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		// Note: The current implementation returns 500 for errors, ideally it should be 400 or 422 for business logic errors.
		// But we are testing existing behavior + our fixes.
		if w.Code != 500 {
			t.Errorf("Expected status 500, got %d", w.Code)
		}
	})
}
