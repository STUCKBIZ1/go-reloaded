package go_reloaded

import (
	"strings"
)

func Punctuations(str string) string {
	runes := []rune(str)
	i := 0

	for i < len(runes) {
		if IsPunctuation(runes[i]) {
			if i > 0 && runes[i-1] == ' ' {
				runes[i-1], runes[i] = runes[i], runes[i-1]
				i = 0
			}
		}
		i++
	}
	c := strings.Fields(string(runes))
	g := strings.Join(c, " ")
	res := []rune(g)
	result := []rune{}
	for j := 0; j < len(res); j++ {
		if j > 0 && (!IsPunctuation(res[j]) && IsPunctuation(res[j-1])) {
			if j != len(res)-1 {
				result = append(result, ' ')
			}
			result = append(result, res[j])
		} else {
			result = append(result, res[j])
		}
	}
	k := []string{}
	l := strings.Split(string(result), " ")
	for _, v := range l {
		if v != "" {
			k = append(k, v)
		}
	}
	m := strings.Join(k, " ")
	return string(m)
}
