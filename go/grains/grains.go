//Package grains finds number of grains in each square of a chess chessboard
package grains

import (
	"errors"
)

//Square returns a number 2 to the power of number -1
func Square(input int) (uint64, error) {

	if input <= 0 || input > 64 {
		return 0, errors.New("Number cannot be zero or negative or greater than 64")
	}

	return 1 << uint64(input-1), nil
}

//Total returns the total of squares from 1 to 64
func Total() uint64 {
	return ^(uint64(0))
}
