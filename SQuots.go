package main

import (
	"strings"
	"unicode"
)

func IsAlpha(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func CheckSymbol(r rune) bool {
	symbols := ",.!?:;)"
	return strings.ContainsRune(symbols, r)
}

func hasClosingQuote(runes []rune) bool {
	for i, r := range runes {
		if r == '\'' {
			if i > 0 && i+1 < len(runes) && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
				continue
			}
			return true
		}
	}
	return false
}

func SQuots(s string) string {
	runes := []rune(s)
	res := []rune{}
	open := -1

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		if r == '\'' {
			if i > 0 && i+1 < len(runes) && IsAlpha(runes[i-1]) && IsAlpha(runes[i+1]) {
				res = append(res, r)
				continue
			}

			if open == -1 {
				open = len(res)
				if len(res) > 0 && !unicode.IsSpace(res[len(res)-1]) && hasClosingQuote(runes[i+1:]) {
					res = append(res, ' ')
					open++
				}
				res = append(res, '\'')
				continue
			}

			if open != -1 {
				content := strings.TrimSpace(string(res[open+1:]))
				res = res[:open+1]
				res = append(res, []rune(content)...)
				res = append(res, '\'')
				if i+1 < len(runes) && !unicode.IsSpace(runes[i+1]) && !CheckSymbol(runes[i+1]) {
					res = append(res, ' ')
				}
				open = -1
				continue
			}
		} else {
			res = append(res, r)
		}
	}

	return string(res)
}
