package summultiples

// SumMultiples find the sum of all the unique multiples
// of particular numbers up to but not including that number
func SumMultiples(limit int, divisors ...int) int {
	var sum int
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}

	return sum
}
