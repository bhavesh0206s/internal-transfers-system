package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/handlers"
)

func SetupTransactionRoutes(router *gin.RouterGroup, h handlers.TransactionsHandler) {
	router.POST("/transactions", h.CreateTransactionHandler)
}
