package go_reloaded

import (
	"strings"
)

func IsVowel(char rune) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
		return true
	}
	return false
}

func SpQ(line string) string {
	words := strings.Split(line, "'")
	var res2 []string
	for _, strs := range words {
		r := strings.Split(strs, " ")
		for i, word := range r {
			if i < len(r)-1 && (word == "a" || word == "A") && IsVowel(rune(r[i+1][0])) {
				r[i] += "n"
			}
		}
		res1 := strings.Join(r, " ")
		res2 = append(res2, res1)
	}
	return strings.Join(res2, "'")
}

func SpS(line string) string {
	words := strings.Split(line, " ")
	for i, word := range words {
		if i < len(words)-1 && (word == "a" || word == "A") && IsVowel(rune(words[i+1][0])) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}

func Vowels(line string) string {
	result1 := SpS(line)
	result2 := SpQ(result1)
	return result2
}
