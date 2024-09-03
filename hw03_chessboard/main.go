package main

import (
	"fmt"
)

func main() {
	var size int

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&size)

	if size == 0 {
		size = 8
	}

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
