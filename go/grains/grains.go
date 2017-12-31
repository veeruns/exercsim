//Package grains finds number of grains in each square of a chess chessboard
package grains

import (
	"errors"
	"math"
)

func Square(input int) (uint64, error) {

	if input <= 0 || input > 64 {
		return 0, errors.New("Number cannot be zero or negative or greater than 64")
	}

	return uint64(math.Pow(float64(2), float64((input - 1)))), nil
}

func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		rep, _ := Square(i)
		sum += rep
	}
	return sum
}
