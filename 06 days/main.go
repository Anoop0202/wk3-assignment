package main

import "fmt"

func main() {
	var day int
	fmt.Println("enter a number between 1 and 7")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("sunday")
		break
	case 2:
		fmt.Println("monday")
		break
	case 3:
		fmt.Println("tuesday")
		break
	case 4:
		fmt.Println("wednesday")
		break
	case 5:
		fmt.Println("thursday")
		break
	case 6:
		fmt.Println("friday")
		break
	case 7:
		fmt.Println("saturday")
		break
	default:
		fmt.Println("invalid input")

	}
}
