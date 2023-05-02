package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("enter the value of a")
	fmt.Scanln(&a)
	fmt.Print("enter the value of b")
	fmt.Scanln(&b)

	//swaping logic
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("after swapping , the value of a is ", a)
	fmt.Println("after swapping ,the value of b is", b)
}
