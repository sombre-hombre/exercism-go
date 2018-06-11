package romannumerals

import (
	"errors"
	"strings"
)

var (
	roman  = []byte{'M', 'D', 'C', 'L', 'X', 'V', 'I'}
	arabic = []int{1000, 500, 100, 50, 10, 5, 1}
)

// ToRomanNumeral converts from normal numbers to Roman Numerals
func ToRomanNumeral(n int) (r string, e error) {
	if n < 1 {
		return "", errors.New("Number is too small")
	}
	if n > 3000 {
		return "", errors.New("Number is too large")
	}

	var sb strings.Builder

	for i := 0; i < len(roman); i += 2 {
		m := n / arabic[i]
		n = n % arabic[i]
		switch {
		case m == 9:
			sb.WriteByte(roman[i])
			sb.WriteByte(roman[i-2])
		case m > 5:
			sb.WriteByte(roman[i-1])
			sb.WriteString(strings.Repeat(string(roman[i]), m-5))
		case m == 5:
			sb.WriteByte(roman[i-1])
		case m > 3:
			sb.WriteByte(roman[i])
			sb.WriteByte(roman[i-1])
		case m > 0:
			sb.WriteString(strings.Repeat(string(roman[i]), m))
		}
	}

	return sb.String(), nil
}
