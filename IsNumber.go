package main

import (
	"strconv"
	"strings"
)

func IsNumber(str string) bool {
	if str[len(str)-1] == ')' {
		str = strings.TrimRight(str, ")")
		_, err := strconv.Atoi(str)
		if err == nil {
			return true
		}
	}
	return false
}
