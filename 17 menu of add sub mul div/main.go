package main

import "fmt"

func main() {
	var choice int
	var num1, num2 float64
	for {
		fmt.Println("enter your choice")
		fmt.Println("1.addition")
		fmt.Println("2.subtraction")
		fmt.Println("3.multiplication")
		fmt.Println("4.division")
		fmt.Println("5.exit")

		fmt.Scanln(&choice)

		if choice == 5 {
			break
		}
		fmt.Println("enter two numbers")
		fmt.Scanln(&num1)
		fmt.Scanln(&num2)

		switch choice {
		case 1:
			fmt.Printf("result %.2f\n\n", num1+num2)
		case 2:
			fmt.Printf("result %.2f\n\n", num1-num2)
		case 3:
			fmt.Printf("result %.2f\n\n", num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Print("\nerror:cannot divide by zero\n")
			} else {
				fmt.Printf("result %.2f\n\n", num1/num2)
			}
		}
	}
}
