// Package twofer is short for two for one. One for you and one for me.
package twofer

import "fmt"

// ShareWith accepts a string as a parameter. If variable name is an empty
// return a string "One for you, one for me". Otherwise, return "One for
// "{name}, one for me"
func ShareWith(name string) string {
	result := "One for %s, one for me."
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf(result, name)
}
