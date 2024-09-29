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

	chessboerd := GetChessboerd(size)
	fmt.Print(chessboerd)
}

func GetChessboerd(size int) string {

	var result string

	isSpace := false
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			if isSpace {
				result = result + " "
			} else {
				result = result + "#"
			}
			isSpace = !isSpace
		}
		if size%2 == 0 {
			isSpace = !isSpace
		}
		result = result + "\r\n"
	}

	return result
}
