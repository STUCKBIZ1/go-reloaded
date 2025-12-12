package go_reloaded

func IsGFlage(str string) bool {
	switch str {
	case "(up,", "(cap,", "(low,":
		return true
	}
	return false
}
