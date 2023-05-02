package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("enter mark percentage")
	fmt.Scan(&mark)
	switch {
	case mark >= 90:
		fmt.Println("A grade")
		break
	case mark >= 80:
		fmt.Println("B grade")
		break
	case mark >= 70:
		fmt.Println("C grade")
		break
	case mark >= 60:
		fmt.Println("D grade")
		break
	case mark >= 50:
		fmt.Println("E grade")
		break
	case mark < 50 && mark > 0:
		fmt.Println("sorry you are failed")
		break
	default:
		fmt.Println("invalid input")

	}
}
