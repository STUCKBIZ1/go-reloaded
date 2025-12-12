package go_reloaded

func IsFlage(word string) bool {
	switch word {
	case "(up)", "(cap)", "(low)", "(bin)", "(hex)":
		return true
	}
	return false
}
