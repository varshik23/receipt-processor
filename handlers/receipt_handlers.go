package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/varshik23/receipt-processor/utils"
)

func GetReceipt(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func PostReceipt(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	id := utils.Hash(newReceipt)
	points[id] = utils.calculatePoints(newReceipt)
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}