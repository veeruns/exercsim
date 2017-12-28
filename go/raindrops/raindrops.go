//Package raindrops depending on factor 3,5,7 returns a string
package raindrops

import (
	// Efficient string concatenation  using bytes
	"bytes"
	// Converting integer to string using Itoa function
	"strconv"
	// we need to sort the keys in the right order
	"sort"
	// debugging
)

//Convert changes factors to raindrops
func Convert(input int) string {
	factors := make(map[int]string)
	factors[3] = "Pling"
	factors[5] = "Plang"
	factors[7] = "Plong"
	sortedKeys := make([]int, 0, len(factors))
	for key := range factors {
		sortedKeys = append(sortedKeys, key)
	}
	// Now they are always sorted
	sort.Ints(sortedKeys)

	var retstring bytes.Buffer

	for _, val := range sortedKeys {

		v := factors[val]
		if input%val == 0 {
			retstring.WriteString(v)
		}
	}
	if retstring.String() == "" {
		retstring.WriteString(strconv.Itoa(input))
	}
	return retstring.String()
}
