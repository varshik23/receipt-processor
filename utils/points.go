package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
	"github.com/varshik23/receipt-processor/models"
)

func CalculatePoints(receipt models.Receipt) float64 {
	var totalPoints float64 = 0.0
	totalPoints += calculateAlphanumericPoints(receipt)
	totalPoints += calculateTotalBasedPoints(receipt)
	totalPoints += calculateItemBasedPoints(receipt)
	totalPoints += calculateDateAndTimePoints(receipt)
	return totalPoints
}

func calculateAlphanumericPoints(receipt models.Receipt) float64{
	count := 0.0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	fmt.Println("Alphanumeric points: ", count)
	return count
}

func calculateTotalBasedPoints(receipt models.Receipt) float64 {
	total, _ := strconv.ParseFloat(receipt.Total, 64)

	var res float64 = 0.0
	if math.Mod(total, 1.0) == 0 {
		res += 50.0
	}
	if math.Mod(total, 0.25) == 0 {
		res += 25.0
	}
	fmt.Println("Total based points: ", res)
	return res
}

func calculateItemBasedPoints(receipt models.Receipt) float64 {
	var res float64 = 0.0
	res += float64(len(receipt.Items) / 2.0) * 5.0
	for _, item := range receipt.Items {
		item.ShortDescription = strings.TrimSpace(item.ShortDescription)
		if len(item.ShortDescription) != 0 && len(item.ShortDescription) % 3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			res += math.Ceil(price * 0.2)
		}
	}
	fmt.Println("Item based points: ", res)
	return res
}

func calculateDateAndTimePoints(receipt models.Receipt) float64 {
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	tme, _ :=  time.Parse("15:04", receipt.PurchaseTime)

	var res float64 = 0.0
	if date.Day() % 2 != 0 {
		res += 6.0
	}

	Hour := tme.Hour()
	Minute := tme.Minute()

	if ((Hour >= 14 && Hour < 16) && (Minute > 0 && Minute <= 59)){
		res += 10.0
	}
	fmt.Println("Date and time points: ", res)

	return res
}
