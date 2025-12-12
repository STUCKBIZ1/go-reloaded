package go_reloaded

import (
	"strconv"
)

func IsNumber(str string) bool {
	if str == "" {
		return false
	}
	if str[len(str)-1] == ')' {
		str = str[:len(str)-1]
		_, err := strconv.Atoi(str)
		if err == nil {
			return true
		}
	}
	return false
}
