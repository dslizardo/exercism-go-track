package diffsquares

// SquareOfSum returns square of the sum of the first n numbers
func SquareOfSum(n int) int {

	var sum = float64(n) / 2 * float64(1+n)
	return int(sum * sum)
}

// SumOfSquares returns sum of the squares of the n numbers is
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
