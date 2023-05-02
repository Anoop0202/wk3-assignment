package main

import "fmt"

func getArray() [][]int {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	return arr
}
func displayArray(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}
func main() {
	arr := getArray()
	displayArray(arr)
}
