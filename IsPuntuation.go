package main

func IsPunctuation(a rune) bool {
	switch a {
	case '.', ',', ';', ':', '?', '!':
		return true
	}
	return false
}
