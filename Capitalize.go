package main

func Capitalize(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + 32)
		} else {
			result += string(s[i])
		}
	}

	final := ""
	for i := 0; i < len(result); i++ {
		if i == 0 && result[i] >= 'a' && result[i] <= 'z' {
			final += string(result[i] - 32)
		} else if i > 0 {
			if !((result[i-1] >= 'a' && result[i-1] <= 'z') || (result[i-1] >= 'A' && result[i-1] <= 'Z') || (result[i-1] >= '0' && result[i-1] <= '9')) && result[i] >= 'a' && result[i] <= 'z' {
				final += string(result[i] - 32)
			} else {
				final += string(result[i])
			}
		} else {
			final += string(result[i])
		}
	}
	return final
}
