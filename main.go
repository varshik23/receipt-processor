package main

import (
	"github.com/gin-gonic/gin"
	"github.com/varshik23/receipt-processor/api"
)

func main() {
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run(":8080")
}
