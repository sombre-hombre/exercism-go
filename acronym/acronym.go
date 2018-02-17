// Package acronym implements the Acronym exercise.
package acronym

import "strings"
import "bytes"
import "unicode"

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	buffer := bytes.NewBufferString("")

	for _, word := range strings.FieldsFunc(s, isSpaceOrPunct) {
		buffer.WriteByte(word[0])
	}

	return strings.ToUpper(buffer.String())
}

func isSpaceOrPunct(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}
