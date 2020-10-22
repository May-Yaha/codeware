package fix_string_case

import (
	"strings"
	"unicode"
)

// https://www.codewars.com/kata/5b180e9fedaa564a7000009a/train/go
func solve(str string) string {
	if str == "" {
		return str
	}

	upperCount := 0
	lowerCount := 0

	for _, s := range str {
		if unicode.IsUpper(s) {
			upperCount++
		} else {
			lowerCount++
		}
	}

	if upperCount > lowerCount {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
