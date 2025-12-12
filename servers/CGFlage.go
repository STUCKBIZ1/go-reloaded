package go_reloaded

import (
	"strconv"
	"strings"
	 
)

func CGFlage(words []string, str string, flag string) []string {
	var strs []string
	for i, v := range words {
		if v != "" && i < len(words) {
			strs = append(strs, v)
		}
	}
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
