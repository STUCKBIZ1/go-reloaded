package main

import (
	"strconv"
	"strings"
)

func CGFlage(strs []string, str string, flag string) []string {
	switch flag {
	case "(cap,":
		str = strings.TrimRight(str, ")")
		num, _ := strconv.Atoi(str)
		i := len(strs) - 1
		for i >= 0 {
			if num > 0 && IsValid(strs[i]) {
				strs[i] = Capitalize(strs[i])
				num--
			}
			i--
		}
		return strs
	case "(up,":
		str = strings.TrimRight(str, ")")
		num, _ := strconv.Atoi(str)
		i := len(strs) - 1
		for i >= 0 {
			if num > 0 && IsValid(strs[i]) {
				strs[i] = strings.ToUpper(strs[i])
				num--
			}
			i--
		}
		return strs
	case "(low,":
		str = strings.TrimRight(str, ")")
		num, _ := strconv.Atoi(str)
		i := len(strs) - 1
		for i >= 0 {
			if num > 0 && IsValid(strs[i]) {
				strs[i] = strings.ToLower(strs[i])
				num--
			}
			i--
		}
		return strs
	}
	return strs
}
