package kata

import "strings"

func CamelCase(s string) (result string) {
	var isUpper = true

	for _, v := range s {
		if v != ' ' {
			char := string(v)
			if isUpper {
				char = strings.ToUpper(char)
				isUpper = false
			}
			result += char
		} else {
			isUpper = true
		}

	}
	return
}
