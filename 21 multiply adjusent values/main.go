package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8}
	result := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		result[i] = arr[i] * arr[i+1]
	}
	fmt.Println(result)
}
