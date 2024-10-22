package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "зима осень  зима лето осень лето зима весна"
	result := CountWords(text)
	fmt.Println(result)
}

func CountWords(s string) map[string]int {
	var LowerWord string

	result := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		LowerWord = strings.ToLower(word)
		result[LowerWord]++
	}
	return result
}
