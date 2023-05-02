package main

import "fmt"

func main() {
	var array1, array2 [10]int
	var n1, n2, temp int
	fmt.Println("enter the limit of the array1")
	fmt.Scan(&n1)
	fmt.Println("enter the values in array1")
	for i := 0; i < n1; i++ {
		fmt.Scan(&array1[i])
	}
	fmt.Println("enter the limit of the array2")
	fmt.Scan(&n2)
	fmt.Println("enter the values in array2 ")
	for i := 0; i < n2; i++ {
		fmt.Scan(&array2[i])
	}
	for i := 0; i < n2; i++ {
		temp = array1[i]
		array1[i] = array2[i]
		array2[i] = temp

	}
	fmt.Print(" array1 \t")
	for i := 0; i < n2; i++ {
		fmt.Print(array1[i])
		fmt.Print("\t")
	}
	fmt.Print("\narray2 \t")
	for i := 0; i < n2; i++ {
		fmt.Print(array2[i])
		fmt.Print("\t")
	}

}
