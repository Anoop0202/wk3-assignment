package main

import "fmt"

func main() {
	var a1 [10][10]int
	a2 := a1
	var n int
	fmt.Println("enter the limit of the array")
	fmt.Scan(&n)
	fmt.Println("enter the elements of first metrix")
	for i := 0; i < n; i++ {
		fmt.Printf("enter the elements of a %v row\n", i)
		for j := 0; j < n; j++ {
			fmt.Scan(&a1[i][j])
		}
	}
	fmt.Println("enter the elements of second metrix")
	for i := 0; i < n; i++ {
		fmt.Printf("enter the elements of a %v row\n", i)
		for j := 0; j < n; j++ {
			fmt.Scan(&a2[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a1[i][j] = a2[i][j] * a1[i][j]
		}
	}
	fmt.Printf("the multiplied metrix is \n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v", a1[i][j])
			fmt.Print("\t")
		}
		fmt.Printf("\n")
	}
}
