package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/varshik23/receipt-processor/models"
)

func TestCalculatePoints1(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	}
	expected := 28.0
	actual := CalculatePoints(receipt)
	assert.Equal(t, expected, actual)
}

func TestCalculatePoints2(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []models.Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			}, {
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
		Total: "9.00",
	}
	expected := 109.0
	actual := CalculatePoints(receipt)
	assert.Equal(t, expected, actual)
}

func TestCalculatePoints3(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Walmart",
		PurchaseDate: "2023-05-15",
		PurchaseTime: "15:45",
		Items: []models.Item{
			{
				ShortDescription: "Coca-Cola 2L",
				Price:            "2.99",
			}, {
				ShortDescription: "Doritos Cool Ranch",
				Price:            "3.50",
			}, {
				ShortDescription: "Kleenex Tissues",
				Price:            "1.75",
			}, {
				ShortDescription: "Colgate Toothpaste",
				Price:            "2.99",
			},
		},
		Total: "11.23",
	}

	expected := 37.0
	actual := CalculatePoints(receipt)
	assert.Equal(t, expected, actual)
}

func TestCalculatePoints4(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Best Buy",
		PurchaseDate: "2023-11-10",
		PurchaseTime: "16:20",
		Items: []models.Item{
			{
				ShortDescription: "Sony Wireless Earbuds",
				Price:            "69.99",
			}, {
				ShortDescription: "Logitech Gaming Mouse",
				Price:            "49.99",
			}, {
				ShortDescription: "Samsung 4K Smart TV",
				Price:            "599.99",
			},
		},
		Total: "719.97",
	}

	expected := 36.0
	actual := CalculatePoints(receipt)
	assert.Equal(t, expected, actual)
}
