package main

import "fmt"

func main() {
	var a [10]int
	var n int
	fmt.Println("enter the limit of the array")
	fmt.Scan(&n)
	fmt.Println("enter the elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Print("the even numbers are ")
	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			fmt.Print(a[i])
			fmt.Print(",")
		}
	}

}
