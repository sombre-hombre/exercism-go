package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	return result * result
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}

	return result
}

// Difference returns the difference between the square of the sum of the first
// n natural numbers and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
