//Package twofer returns a specified based on input
package twofer

import (
	"fmt"
	"regexp"
)

const testVersion = 3

//ShareWith function input is a string and output is a message
func ShareWith(input string) string {
	var output string
	var isEmpty bool
	isEmpty, _ = regexp.MatchString("^\\s*$", input)

	if isEmpty == false {
		output = fmt.Sprintf("One for %s, one for me.", input)
	} else {
		output = fmt.Sprintf("One for you, one for me.")
	}
	return output
}
