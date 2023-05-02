package main

import (
	"fmt"
	"strings"
)

func main() {
	var array string
	var n, i int
	fmt.Println("enter the string")
	fmt.Scan(&array)
	array = strings.ToLower(array)
	isPalindrome := true
	n = len(array) - 1
	for i = 0; i <= n; i++ {
		if array[i] != array[n-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome == true {
		fmt.Println("the string is palindrome")
	} else {
		fmt.Println("the string is not palindrome")
	}

}
