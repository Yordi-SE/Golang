package main

import (
	"regexp"
	"strings"
)

func WordCount(s string) map[string]int {
	re := regexp.MustCompile(`[a-zA-Z]+`)

	words := re.FindAllString(strings.ToLower(s), -1)

	wordFreq := make(map[string]int)

	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}

func IsPalindrome(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)

	chars := re.FindAllString(strings.ToLower(s), -1)

	cleaned := strings.Join(chars, "")

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}
