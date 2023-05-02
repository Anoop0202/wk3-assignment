package main

// program to accept a char value from user and display it on console
import "fmt"

func main() {
	var name string
	fmt.Print("enter your name\n")
	fmt.Scan(&name)
	fmt.Printf("you entered name is %s", name)
}
