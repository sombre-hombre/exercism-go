package sieve

// goos: windows
// goarch: amd64
// BenchmarkSieve-8   	  300000	      4690 ns/op	    9440 B/op	       8 allocs/op

// Sieve of Eratosthenes to find all the primes from 2
// up to a given number
func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	m := make([]bool, limit+1)
	result := make([]int, 0, limit-1)
	step := 2

	for i := 2; i <= limit; i++ {
		if m[i] {
			continue
		}

		if i == 2 {
			step = 1
		}

		for j := i * i; j <= limit; j += step * i {
			m[j] = true
		}
		result = append(result, i)
	}

	return result
}
