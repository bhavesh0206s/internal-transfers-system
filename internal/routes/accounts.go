package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/handlers"
)

func SetupAccountRoutes(router *gin.RouterGroup, h handlers.AccountHandler) {
	router.GET("/accounts/:account_id", h.GetAccountBalanceHandler)
	router.POST("/accounts", h.CreateAccountHandler)
}
