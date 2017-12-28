//Package bob is apparently a bot. Very taciturn
package bob

import (
	"regexp"
)

// Hey function returns based on some rules, if its a question,shouting etc
func Hey(remark string) string {
	// Is this an empty String
	emptyQuestion, _ := regexp.MatchString("^\\s*$", remark)
	//initiazlise regexp

	revar := regexp.MustCompile("[^a-zA-Z?]*")
	// replace all non alpha numberic characters and question mark with nothing

	replacedString := revar.ReplaceAllString(remark, "")

	// is this a question i.e has a question mark in the end

	reQuestion, _ := regexp.MatchString(".*\\?$", replacedString)

	// Is this being yelled and a Question

	reyellquestion, _ := regexp.MatchString("^[A-Z]+\\?$", replacedString)

	// Is this just plain yelling

	reallcaps, _ := regexp.MatchString("^[A-Z]*[^\\?]$", replacedString)

	switch {
	case emptyQuestion:
		return "Fine. Be that way!"
	case reQuestion && reyellquestion:
		return "Calm down, I know what I'm doing!"
	case reQuestion && !reyellquestion:
		return "Sure."
	case !reQuestion && reallcaps:
		return "Whoa, chill out!"
	default:
		return "Whatever."

	}

}
