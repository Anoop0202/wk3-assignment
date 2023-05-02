package main

import "fmt"

func main() {

	var a [10]int
	var n, temp int
	fmt.Println("enter a limit")
	fmt.Scan(&n)

	fmt.Println("enter the values")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] < a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
		fmt.Print(a[i])
		fmt.Print(" > ")
	}

}
