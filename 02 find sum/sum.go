package main

import "fmt"

//program to accept 1 int and 1 float value as input and display the sum as float value
func main() {
	var num1 int
	var num2 float64
	fmt.Printf("enter int value")
	fmt.Scan(&num1)
	fmt.Printf("enter float value")
	fmt.Scan(&num2)
	sum := num2 + float64(num1)
	fmt.Printf("sum is %.2f\n", sum)
}
