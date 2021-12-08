package section3

import (
	"fmt"
	"os"
	"strconv"
)

const (
	EUR = iota
	GBP
	JPY
)
const errMsg = "Please provide the amount to be converted"

func currencyConverter() {
	var cc [3]float64 = [3]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	if len(os.Args) != 2 {
		fmt.Printf("%s", errMsg)
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, cc[EUR]*amount)
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, cc[GBP]*amount)
	fmt.Printf("%.2f USD is %.2f JPY\n", amount, cc[JPY]*amount)
}
