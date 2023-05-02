package main

import "fmt"

func main() {
	var num int

	fmt.Println("enter a number")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v*%v=%v", i, num, num*i)
		fmt.Printf("\n")
	}
}
