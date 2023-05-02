package main

import "fmt"

func main() {
	var annualIncome float64
	var taxAmount float64
	fmt.Println("enter the annual income")
	fmt.Scanln(&annualIncome)
	switch {
	case annualIncome <= 250000:
		taxAmount = 0
	case annualIncome <= 500000:
		taxAmount = (annualIncome - 250000) * 0.05
	case annualIncome <= 1000000:
		taxAmount = 12500 + (annualIncome-500000)*0.2
	default:
		taxAmount = 112500 + (annualIncome-1000000)*0.3
	}
	fmt.Printf("the income tax amount is %.2f", taxAmount)
}
