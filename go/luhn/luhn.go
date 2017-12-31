//Package luhn uses luhn algorithm to validate a credit card numbers
package luhn

import (
	"strconv"
	"strings"
)

// Valid checks if the luhn checksum for a number is valid or not
func Valid(input string) bool {
	nr := strings.NewReplacer(" ", "")
	upperinput := nr.Replace(input)
	converted, err := strconv.Atoi(upperinput)
	if err == nil {

	} else {
		return false
	}
	if len(upperinput) < 2 {
		return false
	}

	var sum int
	for i := 0; i < len(upperinput); i++ {
		residue := converted % 10

		if i%2 != 0 {
			residue = residue * 2
			if residue > 9 {
				residue = residue - 9
			}
		}
		sum = sum + residue

		converted = converted / 10
	}

	if sum%10 == 0 {
		return true
	}
	return false

}
