package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(input string) bool {
	// using map as set
	letters := map[rune]struct{}{}

	for _, r := range strings.ToUpper(input) {
		// skip non letters
		if !unicode.IsLetter(r) {
			continue
		}

		// if set already contains the letter
		if _, found := letters[r]; found {
			// then it's not an isogram
			return false
		}

		// memoise the letter
		letters[r] = struct{}{}
	}

	return true
}
