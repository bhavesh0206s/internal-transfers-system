package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thebhavesh02/internal-transfers-service/internal/db"
	"github.com/thebhavesh02/internal-transfers-service/internal/handlers"
)

func RegisterRoutes(routes *gin.Engine, queries db.Queries) {
	router := routes.Group("/api/v1")

	SetupAccountRoutes(router, handlers.NewAmountHandler(queries))
	SetupTransactionRoutes(router, handlers.NewTransactionsHandler(queries))
}
