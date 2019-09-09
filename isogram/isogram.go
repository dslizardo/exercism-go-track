package isogram

import "unicode"

//IsIsogram determines if the given string is an isogram
func IsIsogram(input string) bool {
	letters := make(map[rune]bool)

	for _, c := range input {
		c = unicode.ToLower(c)

		if c == '-' || c == ' ' {
			continue
		}

		if letters[c] {
			return false
		}

		letters[c] = true
	}

	return true
}
