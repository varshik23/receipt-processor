package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func calculatePoints(receipt receipt) float64 {

	var totalPoints float64 = 0.0
	// totalPoints += len(receipt.Retailer)
	totalPoints += calculateAlphanumericPoints(receipt)

	total, error := strconv.ParseFloat(receipt.Total, 64)

	if error != nil {
		fmt.Println(error)
	}

	if math.Mod(total, 1.0) == 0 {
		totalPoints += 50.0
	}
	fmt.Printf("Total Rounded: %f\n", totalPoints)
	if math.Mod(total, 0.25) == 0 {
		totalPoints += 25.0
	}
	fmt.Printf("Total 0.25: %f\n", totalPoints)
	totalPoints += float64(len(receipt.Items) / 2.0) * 5.0
	fmt.Printf("Total 2 pairs: %f\n", totalPoints)
	for _, item := range receipt.Items {
		item.ShortDescription = strings.TrimSpace(item.ShortDescription)
		if len(item.ShortDescription) % 3 == 0 {
			price, error := strconv.ParseFloat(item.Price, 64)
			if error != nil {
				fmt.Println(error)
			}
			totalPoints += math.Ceil(price * 0.2)
		}
		fmt.Printf("Total Item: %f\n", totalPoints)
	}
	totalPoints += calculateDateAndTimePoints(receipt)
	fmt.Printf("Total points: %f\n", totalPoints)
	return totalPoints
}

func calculateAlphanumericPoints(receipt receipt) float64{
	count := 0.0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}
func calculateDateAndTimePoints(receipt receipt) float64 {
	date, d_error := time.Parse("2006-01-02", receipt.PurchaseDate)
	tme, t_error :=  time.Parse("15:04", receipt.PurchaseTime)

	if d_error != nil {
		fmt.Println(d_error)
	}
	if t_error != nil {
		fmt.Println(t_error)
	}

	var res float64 = 0.0
	if date.Day() % 2 != 0 {
		res += 6.0
	}

	Hour := tme.Hour()
	Minute := tme.Minute()

	if ((Hour >= 14 && Hour < 16) && (Minute > 0 && Minute <= 59)){
		res += 10.0
	}

	return res
}
