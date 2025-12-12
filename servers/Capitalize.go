package go_reloaded

import "unicode"

func Capitalize(s string) string {
	result := ""
	count := 0
	for _, v := range s {
		if unicode.IsLetter(v) && count == 0 {
			count = 1
			result += string(unicode.ToUpper(v))
		} else {
			result += string(unicode.ToLower(v))
		}
	}
	return result
}
