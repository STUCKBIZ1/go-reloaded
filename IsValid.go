package main

func IsValid(word string) bool {
	for _, char := range word {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}
