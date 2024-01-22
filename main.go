package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []item `json:"items"`
	Total string `json:"total"`
}

type item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

var receipts = []receipt{
	{
		Retailer: "Walmart",
		PurchaseDate: "2020-01-01",
		PurchaseTime: "12:00:00",
		Items: []item{
			{
				ShortDescription: "Milk",
				Price: "2.99",
			},
			{
				ShortDescription: "Eggs",
				Price: "1.99",
			},
		},
		Total: "4.98",
	},
}

func main() {
	router := gin.Default()
	router.GET("/receipt", getReceipt)
	router.POST("/receipt", postReceipt)
	router.Run("localhost:8080")
}

func getReceipt(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func postReceipt(c *gin.Context) {
	var newReceipt receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}