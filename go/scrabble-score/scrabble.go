//Package scrabble gives score of a word depending on points table listed
package scrabble

import (
	"strings"
)

//Score outputs score of a given word
func Score(input string) int {
	points := make(map[string]int, 26)
	points["A"] = 1
	points["E"] = 1
	points["I"] = 1
	points["O"] = 1
	points["U"] = 1
	points["L"] = 1
	points["N"] = 1
	points["R"] = 1
	points["S"] = 1

	points["T"] = 1
	//two pointers
	points["D"] = 2
	points["G"] = 2
	// three pointers

	points["B"] = 3
	points["C"] = 3
	points["M"] = 3
	points["P"] = 3
	// four pointers

	points["F"] = 4
	points["H"] = 4
	points["V"] = 4
	points["W"] = 4
	points["Y"] = 4
	//five pointers
	points["K"] = 5
	// eight pointers
	points["J"] = 8
	points["X"] = 8
	// ten pointers
	points["Q"] = 10
	points["Z"] = 10
	upper := strings.ToUpper(input)
	var totalpoints int
	for _, k := range upper {
		totalpoints = totalpoints + points[string(k)]
	}

	return totalpoints
}
