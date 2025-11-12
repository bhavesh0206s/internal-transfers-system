package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/db"
	"github.com/thebhavesh02/internal-transfers-service/internal/pkg"
	"github.com/thebhavesh02/internal-transfers-service/internal/repositories"
)

type AccountHandler interface {
	GetAccountBalanceHandler(c *gin.Context)
	CreateAccountHandler(c *gin.Context)
}

type accountHandler struct {
	accountRespository repositories.AccountRepository
}

func NewAmountHandler(queries db.Queries) AccountHandler {
	return &accountHandler{
		accountRespository: repositories.NewAccountRepository(queries),
	}
}

func (ah *accountHandler) CreateAccountHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	var body pkg.Account

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatalln(err.Error())
		c.JSON(400, gin.H{"status": "failed", "message": "Something went wrong"})
		return
	}
	ctx := c.Request.Context()

	if body.AccountID <= 0 {
		c.JSON(400, gin.H{"status": "failed", "message": "invalid accountID"})
		return
	}

	if body.InitialBalance == "" {
		c.JSON(400, gin.H{"status": "failed", "message": "initialBalance is required"})
		return
	}

	err := ah.accountRespository.CreateAccount(ctx, &body.AccountID, &body.InitialBalance)

	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success", "message": "account created successfully"})
}

func (ah *accountHandler) GetAccountBalanceHandler(c *gin.Context) {
	accountIDStr := c.Param("account_id")
	if accountIDStr == "" {
		c.JSON(400, gin.H{"status": "failed", "message": "account_id query parameter is required"})
		return
	}

	accountIDInt, err := strconv.ParseInt(accountIDStr, 10, 64)

	if err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": "invalid account_id"})
		return
	}

	ctx := c.Request.Context()
	accountID, balanceId, err := ah.accountRespository.GetAccountBalance(ctx, accountIDInt)

	if err != nil {
		c.JSON(500, gin.H{"status": "failed", "message": "No account found with given account id"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": map[string]interface{}{"accountID": accountID, "balance": balanceId}})
}
