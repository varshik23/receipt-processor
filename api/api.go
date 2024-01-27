package api

import (
	"github.com/gin-gonic/gin"
	"github.com/varshik23/receipt-processor/handlers"
)

func SetupRoutes(router *gin.Engine)  {
	router.GET("/receipt", handlers.GetReceipts)
	router.GET("/receipt/:id/points", handlers.GetReceiptById)
	router.POST("/receipt", handlers.PostReceipt)
	router.DELETE("/receipt/:id", handlers.DeleteReceipt)
}
