package go_reloaded

func IsPunctuation(a rune) bool {
	switch a {
	case '.', ',', ';', ':', '?', '!':
		return true
	}
	return false
}
