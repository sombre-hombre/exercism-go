package perfect

import (
	"errors"
	"math"
)

// Classification represent Nicomachus' (60 - 120 CE) classification scheme for natural numbers.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

// ErrOnlyPositive indicates that the input is not a positive integer
var ErrOnlyPositive = errors.New("The input is not a positive integer")

// Classify determine if a number is perfect, abundant,
// or deficient based on Nicomachus' (60 - 120 CE) classification
// scheme for natural numbers.
// BenchmarkClassify-8   	   10000	    116705 ns/op	       0 B/op	       0 allocs/op
func Classify(n int64) (c Classification, e error) {
	if n <= 0 {
		return c, ErrOnlyPositive
	}

	var s int64 = 1
	var r = int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= r; i++ {
		if n%i == 0 {
			s += i

			ni := n / i
			if i != ni {
				s += ni
			}

			if s > n {
				break
			}
		}
	}

	switch {
	case s < n || n == 1:
		return ClassificationDeficient, nil
	case s == n:
		return ClassificationPerfect, nil
	default:
		return ClassificationAbundant, nil
	}
}
