package lsproduct

import (
	"errors"
	"fmt"
	"strings"
)

// LargestSeriesProduct calculate the largest product for a
// contiguous substring of digits of length n.
func LargestSeriesProduct(digits string, span int) (int, error) {
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
	//prev := best
	current := 0
	for i := 1; i <= len(digits)-span; i++ {
		//first := toInt(digits[i])
		last := toInt(digits[i+span-1])
		if last == 0 {
			i += span - 1
			continue
		}
		//current = (prev / toInt(digits[i-1])) * toInt(digits[i+span])
		current = product(digits[i : i+span])
		if best < current {
			best = current
		}
		//prev = current
	}

	return best, nil
}

func toInt(r byte) int {
	return int(r - '0')
}

func product(wnd string) int {
	if strings.Contains(wnd, "0") {
		return 0
	}

	s := 1
	for _, d := range wnd {
		s *= int(d - '0')
	}

	return s
}
