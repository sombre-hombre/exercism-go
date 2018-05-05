package sieve

// goos: windows
// goarch: amd64
// BenchmarkSieve-8   	   10000	    115100 ns/op	   66584 B/op	       8 allocs/op

// Sieve of Eratosthenes to find all the primes from 2
// up to a given number
func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	m := make(map[int]interface{}, limit-1)

	for i := 2; i*i <= limit; i++ {
		if _, found := m[i]; found {
			continue
		}

		if i == 2 {
			for j := i * i; j <= limit; j += i {
				if _, found := m[j]; !found {
					m[j] = struct{}{}
				}
			}
			continue
		}

		for j := i * i; j <= limit; j += 2 * i {
			if _, found := m[j]; !found {
				m[j] = struct{}{}
			}
		}
	}

	result := make([]int, 0, limit-1)
	for i := 2; i <= limit; i++ {
		if _, found := m[i]; found {
			continue
		}
		result = append(result, i)
	}

	return result
}
