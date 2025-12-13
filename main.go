package main

import (
	"fmt"
	"strings"
)

func main() {
	var wordList []string
	var cleanText string
	input := "Hola mundo, hola Go! Mundo de Go"
	replacer := strings.NewReplacer(",", "", "!", "")
	cleanText = replacer.Replace(input)
	//for _, v := range input {
	//	//fmt.Println(string(v))
	//	word := string(v)
	//	if strings.ContainsAny(word, ",.!?;:") {
	//		cleanText = strings.Replace(input, string(v), "", -1)
	//		continue
	//	}
	//}

	lower := strings.ToLower(cleanText)
	wordList = strings.Fields(lower)
	result := make(map[string]int)

	for _, word := range wordList {
		result[word]++
		continue
	}

	for key, value := range result {
		fmt.Println(key, value)
	}
}
