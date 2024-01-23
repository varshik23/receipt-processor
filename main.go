package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/varshik23/receipt-processor/api"
)

var points = make(map[string]float64)

var receipts = []receipt{}

func main() {
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run("localhost:8080")
}
