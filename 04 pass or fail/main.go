package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("enter your mark")
	fmt.Scan(&mark)
	if mark < 50 {
		fmt.Println("failed")
	} else {
		fmt.Println("passed")
	}
}
