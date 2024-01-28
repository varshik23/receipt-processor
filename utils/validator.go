package utils

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/varshik23/receipt-processor/models"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "numeric":
		return "Should be a number"
	}
	return "Unknown error"
}

func Validate(receipt models.Receipt) error {
	purchaseDate := receipt.PurchaseDate
	purchaseTime := receipt.PurchaseTime
	items := receipt.Items
	_, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return errors.New("invalid date format")
	}

	_, err = time.Parse("15:04", purchaseTime)
	if err != nil {
		return errors.New("invalid time format")
	}

	if len(items) == 0 {
		return errors.New("items cannot be empty")
	}

	return nil
}
