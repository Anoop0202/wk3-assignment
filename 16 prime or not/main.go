package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter the number")
	fmt.Scan(&n)
	isprime := true
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			isprime = false
		}
	}
	if isprime == true {
		fmt.Printf("prime number")
	} else {
		fmt.Printf("composite number")
	}
}
