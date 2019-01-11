package diffsquares

// SquareOfSum sum values using AP(arithmetic progression) and return the square value
func SquareOfSum(value int) int {
	x := (1 + value) * value / 2
	return x * x
}

// SumOfSquares sum square values using AP(arithmetic progression) and return the value
func SumOfSquares(value int) int {
	x := (1 + value) * (2*value + 1) * value / 6
	return x
}

// Difference subtract the difference between the square of sum and the sum of square
func Difference(value int) int {
	return SquareOfSum(value) - SumOfSquares(value)
}
