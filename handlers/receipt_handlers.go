package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/varshik23/receipt-processor/models"
	"github.com/varshik23/receipt-processor/utils"
)

var receipts = make(map[string]models.Receipt)
var points = make(map[string]float64)

func GetReceipts(c *gin.Context) {
	var receiptsArray []models.Receipt
	for _, value := range receipts {
		receiptsArray = append(receiptsArray, value)
	}
	if len(receiptsArray) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipts found"})
		return
	}
	c.IndentedJSON(http.StatusOK, receiptsArray)
}

func GetReceiptById(c *gin.Context) {
	id := c.Param("id")
	if _, ok := points[id]; ok {
		c.IndentedJSON(http.StatusOK, gin.H{"points": points[id]})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Receipt with id %s not found", id)})
}

func PostReceipt(c *gin.Context) {
	var newReceipt models.Receipt

	var v validator.ValidationErrors
	if err := c.BindJSON(&newReceipt); err != nil {
		if errors.As(err, &v) {
			out := make([]models.ErrorMsg, len(v))
			for i, fe := range v {
				out[i] = models.ErrorMsg{Field: fe.Field(), Message: utils.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	if err := utils.Validate(newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := utils.Hash(newReceipt)

	points[id] = utils.CalculatePoints(newReceipt)
	receipts[id] = newReceipt
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func DeleteReceipt(c *gin.Context) {
	id := c.Param("id")
	if _, ok := points[id]; ok {
		delete(points, id)
		delete(receipts, id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Receipt with id %s deleted", id)})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Receipt with id %s not found", id)})
}
