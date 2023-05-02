package main

import "fmt"

func main() {
	var limit int
	sum := 0

	fmt.Println("enter a limit")
	fmt.Scan(&limit)
	for i := 1; i <= limit; i++ {
		if i%2 == 1 {
			sum = sum + i
		}
	}
	fmt.Printf(" sum of the odd numbers is %v ", sum)
}
