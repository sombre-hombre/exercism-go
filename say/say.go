package say

import (
	"fmt"
	"io"
	"strings"
)

var (
	numbers = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}

	tens = []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	places = []struct {
		value int64
		name  string
	}{
		{1e9, "billion"},
		{1e6, "million"},
		{1e3, "thousand"},
	}
)

func writeTens(n int, wr io.Writer) {
	if n < 20 {
		io.WriteString(wr, numbers[n])
		return
	}

	m := n / 10
	n = n % 10

	if n == 0 {
		io.WriteString(wr, tens[m-1])
		return
	}

	fmt.Fprintf(wr, "%s-%s", tens[m-1], numbers[n])
}

func writeHundreds(n int, wr io.Writer) {
	if n < 100 {
		writeTens(n, wr)
		return
	}

	hundreds := n / 100
	fmt.Fprintf(wr, "%s hundred", numbers[hundreds])
	n = n % 100
	if n > 0 {
		io.WriteString(wr, " ")
		writeTens(n, wr)
	}
}

// Say spell out a number in English.
func Say(n int64) (string, bool) {
	if n < 0 || n >= 1e12 {
		return "", false
	}

	var sb strings.Builder
	if n < 100 {
		writeHundreds(int(n), &sb)
		return sb.String(), true
	}

	for _, p := range places {
		if n < p.value {
			continue
		}
		writeHundreds(int(n/p.value), &sb)
		fmt.Fprintf(&sb, " %s ", p.name)
		n = n % p.value
		if n == 0 {
			break
		}
	}

	if n > 0 {
		writeHundreds(int(n), &sb)
	}

	return strings.TrimSpace(sb.String()), true
}
