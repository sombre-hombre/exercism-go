package pangram

import (
	"strings"
)

// IsPangram determines if an input string is a pangram
func IsPangram(input string) bool {
	// using map as a set
	letters := map[rune]struct{}{}

	for _, r := range strings.ToLower(input) {
		// skip non letters
		if r >= 'a' && r <= 'z' {
			// memoise the letter
			letters[r] = struct{}{}
		}
	}

	return len(letters) > int('z'-'a')
}
