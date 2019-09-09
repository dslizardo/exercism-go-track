// Package raindrops is a variation of famous
// FizzBuzz coding problem
package raindrops

import (
	"strconv"
)

// Convert a number to a string, the contents of
// which depend on the number's factors.
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor,
// just pass the number's digits straight through.
func Convert(n int) string {
	var str string

	if n%3 == 0 {
		str += "Pling"
	}

	if n%5 == 0 {
		str += "Plang"
	}

	if n%7 == 0 {
		str += "Plong"
	}

	if len(str) == 0 {
		str += strconv.Itoa(n)
	}

	return str
}
