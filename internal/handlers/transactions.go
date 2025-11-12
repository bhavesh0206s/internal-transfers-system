package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/db"
	"github.com/thebhavesh02/internal-transfers-service/internal/pkg"
	"github.com/thebhavesh02/internal-transfers-service/internal/repositories"
)

type TransactionsHandler interface {
	CreateTransactionHandler(c *gin.Context)
}

type transactionsHandler struct {
	transactionRespository repositories.TransactionsRepository
}

func NewTransactionsHandler(queries db.Queries) TransactionsHandler {
	return &transactionsHandler{
		transactionRespository: repositories.NewTransactionRepository(queries),
	}
}

func (th *transactionsHandler) CreateTransactionHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var body pkg.TransactionsService

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "Something went wrong"})
		return
	}

	if body.SourceAccountID <= 0 {
		c.JSON(400, gin.H{"status": "failed", "message": "invalid source_account_id"})
		return
	}

	if body.DestinationAccountID <= 0 {
		c.JSON(400, gin.H{"status": "failed", "message": "invalid destination_account_id"})
		return
	}

	if body.SourceAccountID == body.DestinationAccountID {
		c.JSON(400, gin.H{"status": "failed", "message": "source and destination account IDs cannot be the same"})
		return
	}

	amountFloat, err := strconv.ParseFloat(body.Amount, 64)
	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "invalid amount format"})
		return
	}
	if amountFloat <= 0 {
		c.JSON(400, gin.H{"status": "failed", "message": "amount must be greater than zero"})
		return
	}

	ctx := c.Request.Context()

	amountStr := strconv.FormatFloat(amountFloat, 'f', 2, 64)

	err = th.transactionRespository.CreateTransaction(ctx, &body.SourceAccountID, &body.DestinationAccountID, &amountStr)
	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "transaction created"})

}
