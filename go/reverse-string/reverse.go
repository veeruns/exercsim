//Package reverse reverses a strings
package reverse

import "bytes"

const testVersion = 1

//String reverses string
func String(phrase string) string {
	var retstring bytes.Buffer
	runestring := []rune(phrase)
	for i := len(runestring) - 1; i >= 0; i-- {
		retstring.WriteString(string(runestring[i]))
	}
	return retstring.String()
}
