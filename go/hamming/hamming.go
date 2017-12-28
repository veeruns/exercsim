//Package hamming finds hamming distance between two strings returns hamming distance and error object if any
package hamming

import (
	"errors"
)

//Distance function finds the hamming distance.
func Distance(string1, string2 string) (int, error) {

	if len(string1) != len(string2) {
		return -1, errors.New("Input string must match")

	}
	var hammingdistance int
	for i := 0; i < len(string1); i++ {
		if string1[i] != string2[i] {
			hammingdistance++
		}
	}
	return hammingdistance, nil

}
