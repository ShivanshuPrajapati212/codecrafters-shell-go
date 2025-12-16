package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Get all arguments (includes the program path at index 0)
	allArgs := os.Args

	// 2. Print the entire slice at once
	fmt.Println("All arguments as a slice:", allArgs)

	fmt.Println("\nIterating through arguments:")
	// 3. Loop through them to print individually
	// We skip index 0 (the program name) using os.Args[1:]
	for i, arg := range os.Args[1:] {
		fmt.Printf("Argument %d: %s\n", i+1, arg)
	}
}
