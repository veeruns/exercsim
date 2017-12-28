//Package diffsquares has three functions main being find the difference between square of sums and sum of squares
package diffsquares

const testVersion = 1

//SquareOfSums returs square of sum of numbers
func SquareOfSums(input int) int {
	var retval int
	for i := 0; i <= input; i++ {
		retval = retval + i

	}

	return retval * retval
}

//SumOfSquares returns sum of squares of number
func SumOfSquares(input int) int {
	var retval int
	for i := 0; i <= input; i++ {
		retval = retval + i*i
	}
	return retval
}

//Difference returns the difference between square of sums and sum of squares
func Difference(input int) int {
	var retval int
	retval = SquareOfSums(input) - SumOfSquares(input)
	return retval
}
