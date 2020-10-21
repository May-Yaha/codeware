package duplicate_encoder

import (
	"strings"
)

// https://www.codewars.com/kata/54b42f9314d9229fd6000d9c/train/go
func DuplicateEncode(word string) string {

	if word == "" {
		return word
	}

	word = strings.ToLower(word)

	strCountMap := make(map[byte]int)
	for i := range word {
		strCountMap[word[i]]++
	}

	str := ""
	for i := range word {
		if strCountMap[word[i]] > 1 {
			str += ")"
		} else {
			str += "("
		}
	}

	return str
}
