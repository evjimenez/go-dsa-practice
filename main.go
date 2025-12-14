package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hola mundo, hola Go! Mundo de Go"
	result := wordFrequency(text)
	for k, v := range result {
		fmt.Println(k, v)
	}
}

func wordFrequency(txt string) map[string]int {
	result := make(map[string]int)
	lower := strings.ToLower(txt)
	runes := []rune(lower)

	//i = indice y r = rune
	for i, r := range runes {

		if !('a' <= r && r <= 'z') {
			runes[i] = ' '
			continue
		}

	}
	wordList := strings.Fields(string(runes))

	for _, word := range wordList {
		result[word]++
	}

	return result
}
