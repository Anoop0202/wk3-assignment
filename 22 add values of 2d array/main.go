package main

import "fmt"

func createArray(rows, cols int) [][]int {
	arr := make([][]int, rows)
	for i := range arr {
		arr[i] = make([]int, cols)
	}
	return arr
}
func addArrays(arr1, arr2 [][]int) [][]int {
	rows := len(arr1)
	cols := len(arr1[0])
	result := createArray(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	return result
}
func main() {
	arr1 := [][]int{{1, 2}, {3, 4}}
	arr2 := [][]int{{5, 6}, {7, 8}}
	result := addArrays(arr1, arr2)
	fmt.Println(result)
}
