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

	isSpace := false
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if isSpace {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
			isSpace = !isSpace
		}
		if size%2 == 0 {
			isSpace = !isSpace
		}
		fmt.Println()
	}
}
