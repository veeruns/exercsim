//Package isogram finds if a word is an isogram
package isogram

import (
	"strings"
)

//IsIsogram returns a boolean depends on word is Isogram or not
func IsIsogram(input string) bool {

	nr := strings.NewReplacer("-", "", ",", "", " ", "")
	upperinput := nr.Replace(strings.ToUpper(input))

	assoc := make(map[string]int, len(input))
	for _, v := range upperinput {
		value := string(v)
		assoc[value] = assoc[value] + 1
	}
	//	fmt.Printf("%v\n", assoc)
	var retval bool
	// can use to answer how many letters repeated if we want to extend this
	if len(assoc) == len(upperinput) {
		retval = true
	} else {
		retval = false
	}
	return retval
}
