package go_reloaded

import "unicode"

func IsValid(word string) bool {
	for _, char := range word {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}
