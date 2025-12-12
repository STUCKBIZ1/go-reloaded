package go_reloaded

import (
	"strings"
)

func Manager(lines []string) string {
	var result string
	var res []string
	// loop solving line by line
	for _, line := range lines {
		E1 := Flages(line) // for solve the group flages and flages
		E2 := Punctuations(E1)
		E3 := SQuots(E2)
		E4 := Vowels(E3)
		res = append(res, E4)

	}
	result = strings.Join(res, "\n")
	return result
}
