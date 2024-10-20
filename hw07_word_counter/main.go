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
	fmt.Println(s)
	words := strings.Split(s, " ")
	fmt.Println(words)
	for _, word := range words {
		if word != "" {
			LowerWord = strings.ToLower(word)
			result[LowerWord]++
		}
	}
	return result
}
