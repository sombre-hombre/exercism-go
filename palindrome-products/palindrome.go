package palindrome

import (
	"errors"
	"math"
	"strconv"
)

// Product represents result of Products function
type Product struct {
	Product        int
	Factorizations [][2]int
}

func (p *Product) set(product, n, m int) {
	if p.Product != product {
		p.Product = product
		p.Factorizations = [][2]int{[2]int{n, m}}
	} else {
		p.Factorizations = append(p.Factorizations, [2]int{n, m})
	}
}

// Products detects palindrome products in a given range.
// BenchmarkPalindromeProducts-8   	     200	   5996001 ns/op	    2688 B/op	     158 allocs/op
func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var pmin, pmax Product

	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			p := i * j
			if isPalindrome(p) {
				if pmin.Product == 0 || pmin.Product >= p {
					pmin.set(p, i, j)
				}
				if pmax.Product <= p {
					pmax.set(p, i, j)
				}
			}
		}
	}

	if pmin.Product == 0 && pmax.Product == 0 {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return pmin, pmax, nil
}

// BenchmarkIsPalindrome-8   	 1000000	      1110 ns/op	       0 B/op	       0 allocs/op
func isPalindrome1(n int) bool {
	if n < 0 {
		n = -n
	}
	if n < 10 {
		return true
	}

	len := int(math.Floor(math.Log10(float64(n)) + 1))

	for i := 0; i < len/2; i++ {
		little := n / int(math.Pow10(i)) % 10
		big := n / int(math.Pow10(len-i-1)) % 10

		if little != big {
			return false
		}
	}

	return true
}

// BenchmarkIsPalindrome-8   	 3000000	       539 ns/op	      64 B/op	      12 allocs/op
func isPalindrome2(n int) bool {
	s := strconv.Itoa(n)
	if n < 0 {
		s = s[1:]
	}
	len := len(s)

	if len < 2 {
		return true
	}

	for i := 0; i < len/2; i++ {
		if s[i] != s[len-i-1] {
			return false
		}
	}

	return true
}

// BenchmarkIsPalindrome-8   	10000000	       166 ns/op	       0 B/op	       0 allocs/op
func isPalindrome(n int) bool {
	if n < 0 {
		n = -n
	}

	reverse := 0
	copy := n

	for copy != 0 {
		reverse = reverse*10 + copy%10
		copy /= 10
	}

	return n == reverse
}
