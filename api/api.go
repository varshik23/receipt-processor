package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/varshik23/receipt-processor/handlers"
)

func setupRoutes(router *gin.Engine)  {
	router.GET("/receipt", handlers.GetReceipt)
	router.POST("/receipt", handlers.PostReceipt)
}