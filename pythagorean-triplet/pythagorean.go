package pythagorean

import "sort"

// Triplet — pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)
	primitives := getPrimitiveTriplet(0, max)

	for _, t := range primitives {
		if t[0] >= min {
			result = append(result, t)
		}

		for i := 2; ; i++ {
			mt := Triplet{t[0] * i, t[1] * i, t[2] * i}
			if mt[2] > max {
				break
			}
			if mt[0] < min {
				continue
			}
			result = append(result, mt)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

// isCoprime determines is two given numbers are coprime
func isCoprime(x, y int) bool {
	return getGDC(x, y) == 1
}

// getGDC returns the GDC of two given numbers
func getGDC(x, y int) int {
	if x < y {
		x, y = y, x
	}

	if x%y == 0 {
		return y
	}

	x, y = y%x, x-y%x
	return getGDC(x, y)
}

func isEven(n int) bool {
	return n%2 == 0
}

// getPrimitiveTriplet generates primitive Pythagorean triples
// with sides in the range min to max inclusive
func getPrimitiveTriplet(min, max int) []Triplet {
	result := make([]Triplet, 0)

	for n := 1; ; n++ {
		m := n + 1

		if m*m+n*n > max {
			break
		}

		for ; ; m++ {
			if isEven(n-m) || !isCoprime(m, n) {
				// m - n should be odd and m and n should be coprime
				continue
			}

			c := m*m + n*n
			if c > max {
				break
			}
			a := m*m - n*n
			b := 2 * m * n

			if !isEven(n) && !isEven(m) {
				a, b, c = a/2, b/2, c/2
			}

			if a < min || b < min {
				continue
			}

			if a > b {
				a, b = b, a
			}

			result = append(result, Triplet{a, b, c})
		}
	}

	return result
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	result := make([]Triplet, 0)

	for _, t := range Range(0, p) {
		if t[0]+t[1]+t[2] == p {
			result = append(result, t)
		}
	}

	return result
}
