package reverse

import (
	"bytes"
)

// String Reverse a string
func String(input string) string {
	buffer := bytes.NewBufferString("")
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		buffer.WriteRune(runes[i])
	}

	return buffer.String()
}
