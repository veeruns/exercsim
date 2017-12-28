// Package acronym need more test cases
package acronym

// need regexp and strings package to do the trick
import (
	"regexp"
	"strings"
)

// this seems to be a convention followed
const testVersion = 1

// Abbreviate function takes in a string gets first alphabet from each word and converts to uppercase
func Abbreviate(s string) string {
	splitwords := regexp.MustCompile("(\\s+|-)").Split(s, -1)
	output := ""
	for _, value := range splitwords {
		//first alphabet  of the value
		output += string(value[0])

	}
	//convert output to uppercase string
	return strings.ToUpper(output)
}
