// Package hamming implements "Hamming in Go" exercise
package hamming

import (
	"errors"
)

// Distance calculates the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("a and b should have same length")
	}

	result := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			result++
		}
	}

	return result, nil
}
