package main

import "fmt"

func getArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr

}
func display(arr []int, n int) {
	fmt.Printf("the elements of array is \n")
	for i := 0; i < n; i++ {
		fmt.Printf("%v", arr[i])
		fmt.Print(",")
	}
}
func main() {
	var n int
	fmt.Printf("enter the limit of array \n")
	fmt.Scan(&n)
	arr := getArray(n)
	display(arr, n)
}
