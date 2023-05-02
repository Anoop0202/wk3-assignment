package main

import "fmt"

/* finding simple interest using principle amout as ineger value and number of years and rate of interst as float value accept from user ,
calculate simple interest and display it on console - simple interest as the value of float and the specific decimal point is 2 declared in
printf function*/

func main() {
	var P int
	var n, R float64
	fmt.Printf("enter the Principle Amount")
	fmt.Scan(&P)
	fmt.Printf("enter the Number of Years")
	fmt.Scan(&n)
	fmt.Printf("enter the Interest Rate")
	fmt.Scan(&R)
	SimpleInterest := (float64(P) * n * R) / 100
	fmt.Printf("Simple Interest is %.2f\n", SimpleInterest)
}
