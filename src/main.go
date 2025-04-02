package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hi there!\n")

	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("You entered: %d\n", n)
}
