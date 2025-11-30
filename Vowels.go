package main

import "strings"

func IsVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	}
	return false
}

func Vowels(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if i < len(words) - 1 && word == "a" && IsVowel(rune(words[i+1][0]))  || word == "A" && IsVowel(rune(words[i+1][0])) {
			words[i] += "n"
			continue
		}
	}
	result := strings.Join(words, " ")
	return result
}
