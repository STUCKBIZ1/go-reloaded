package main

import (
	"strconv"
	"strings"
)

func CFlage(strs []string, flag string) []string {
	switch flag {
	case "(cap)":
		i := len(strs) - 1
		for i >= 0 {
			if IsValid(strs[i]) {
				strs[i] = Capitalize(strs[i])
				return strs
			} else {
				i--
			}
		}
		return strs
	case "(up)":
		i := len(strs) - 1
		for i >= 0 {
			if IsValid(strs[i]) {
				strs[i] = strings.ToUpper(strs[i])
				return strs
			} else {
				i--
			}
		}
		return strs
	case "(low)":
		i := len(strs) - 1
		for i >= 0 {
			if IsValid(strs[i]) {
				strs[i] = strings.ToLower(strs[i])
				return strs
			} else {
				i--
			}
		}
		return strs
	case "(bin)":
		i := len(strs) - 1
		for i >= 0 {
			dic, err := strconv.ParseInt(strs[i], 2, 64)
			if err == nil {
				strs[i] = strconv.Itoa(int(dic))
				return strs
			} else {
				i--
			}
		}
		return strs
	case "(hex)":
		i := len(strs) - 1
		for i >= 0 {
			dic, err := strconv.ParseInt(strs[i], 16, 64)
			if err == nil {
				strs[i] = strconv.Itoa(int(dic))
				return strs
			} else {
				i--
			}
		}
		return strs
	}
	return strs
}
