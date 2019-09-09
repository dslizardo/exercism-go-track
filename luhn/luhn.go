package luhn

import (
	"strings"
	"unicode"
)

// Valid check if the number passs the Luhn formula
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	sum := 0
	adder := 0
	isSecond := false
	for i := len(input) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(input[i])) {
			return false
		}
		adder = int(input[i] - '0')
		if isSecond {
			adder = (adder * 2)
			if adder > 9 {
				adder = adder - 9
			}
		}
		isSecond = !isSecond
		sum += adder
	}
	return sum%10 == 0
}
