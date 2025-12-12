package go_reloaded

import (
	"strings"
)

func Flages(line string) string {
	var words []string
	var res []string
	var result string
	words = strings.Split(line, " ")
	for i, word := range words {
		if IsFlage(word) {
			if i != 0 {
				switch word {
				case "(cap)", "(up)", "(low)", "(bin)", "(hex)":
					res1 := CFlage(res, word)
					res = res1
				}
			}
		} else if i < len(words)-1 && IsGFlage(word) && IsNumber(words[i+1]) {
			if i != 0 {
				switch word {
				case "(up,", "(cap,", "(low,":
					res1 := CGFlage(res, words[i+1], word)
					res = res1
				}
			}
		} else if i > 0 && IsNumber(word) && IsGFlage(words[i-1]) {
			continue
		} else {
			res = append(res, word)
		}
	}
	result = strings.Join(res, " ")
	return result
}
