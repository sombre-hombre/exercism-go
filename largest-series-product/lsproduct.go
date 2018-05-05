package lsproduct

import (
	"errors"
	"fmt"
	"strings"
)

// goos: windows
// goarch: amd64
// BenchmarkLargestSeriesProduct-8   	 1000000	      1702 ns/op	     104 B/op	       6 allocs/op
// PASS

// LargestSeriesProduct calculate the largest product for a
// contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("String is too short")
	}

	if span < 0 {
		return 0, errors.New("Invalid span size")
	}

	for i, d := range digits {
		if d < '0' || d > '9' {
			return 0, fmt.Errorf("Invalid digit at position %d", i)
		}
	}

	best := product(digits[:span])
	prevProd := best
	currentProd := int64(0)
	for i := 1; i <= len(digits)-span; i++ {
		// last digit in current window
		lastDigit := toInt(digits[i+span-1])
		// last digit before current window
		prevDigit := toInt(digits[i-1])

		if lastDigit == 0 {
			// move window after zero
			i += span - 1
			continue
		}

		if prevDigit == 0 {
			currentProd = product(digits[i : i+span])
		} else {
			currentProd = (prevProd / int64(prevDigit)) * int64(lastDigit)
		}

		if best < currentProd {
			best = currentProd
		}
		prevProd = currentProd
	}

	return best, nil
}

func toInt(r byte) int {
	return int(r - '0')
}

func product(wnd string) int64 {
	if strings.Contains(wnd, "0") {
		return 0
	}

	s := int64(1)
	for _, d := range wnd {
		s *= int64(d - '0')
	}

	return s
}
