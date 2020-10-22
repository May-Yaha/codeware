package alphabetical_addition

// https://www.codewars.com/kata/5d50e3914861a500121e1958/train/go
func AddLetters(letters []rune) rune {
	if len(letters) <= 0 {
		return 'z'
	}

	var count = 0
	for _, letter := range letters {
		count += int(letter) - 96
	}

	c := (count-1)%26 + 97
	return rune(c)
}
