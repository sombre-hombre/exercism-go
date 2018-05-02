package cryptosquare

import (
	"bytes"
	"math"
	"strings"
	"unicode"
)

// Encode encodes input string
// goos: windows
// goarch: amd64
// BenchmarkEncode-8   	  100000	     13343 ns/op	    5824 B/op	      92 allocs/op
func Encode(input string) string {
	normalized := []rune(normalize(input))
	l := len(normalized)
	r, c := getSize(len(normalized))

	var buf bytes.Buffer
	var letter rune
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			idx := j*c + i
			if idx >= l {
				letter = ' '
			} else {
				letter = normalized[idx]
			}

			buf.WriteRune(letter)
		}
		if i < c-1 {
			buf.WriteByte(' ')
		}
	}

	return buf.String()
}

func normalize(input string) string {
	var buf bytes.Buffer
	for _, r := range strings.ToLower(input) {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}
		buf.WriteRune(r)
	}

	return buf.String()
}

// getSize finds out size of a rectangle r×c such that
// `c >= r` and `c - r <= 1`, where `c` is the number of columns
// and `r` is the number of rows.
func getSize(n int) (r, c int) {
	if n == 0 {
		return
	}

	root := math.Sqrt(float64(n))
	r = int(root)
	if n/r == r && n%r == 0 {
		c = r
		return
	}

	c = r + 1
	if r*c < n {
		r++
	}

	return
}
