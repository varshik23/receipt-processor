package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/varshik23/receipt-processor/models"
	"github.com/varshik23/receipt-processor/utils"
)

var receipts = make(map[string]models.Receipt)
var points = make(map[string]float64)

// Handler to get all receipts
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

// Handler to get receipt by id
func GetReceiptById(c *gin.Context) {
	id := c.Param("id")
	if _, ok := points[id]; ok {
		c.IndentedJSON(http.StatusOK, gin.H{"points": points[id]})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Receipt with id %s not found", id)})
}

// Handler to post receipt
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

	id := uuid.New().String()

	points[id] = utils.CalculatePoints(newReceipt)
	receipts[id] = newReceipt
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// Handler to update receipt
func UpdateReceipt(c *gin.Context) {
	id := c.Param("id")
	var updatedReceipt models.Receipt

	var v validator.ValidationErrors
	if err := c.BindJSON(&updatedReceipt); err != nil {
		if errors.As(err, &v) {
			out := make([]models.ErrorMsg, len(v))
			for i, fe := range v {
				out[i] = models.ErrorMsg{Field: fe.Field(), Message: utils.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	if err := utils.Validate(updatedReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if _, ok := points[id]; ok {
		points[id] = utils.CalculatePoints(updatedReceipt)
		receipts[id] = updatedReceipt
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Receipt with id %s updated", id), "updated points": points[id]})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Receipt with id %s not found", id)})
}

// Handler to delete receipt
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
