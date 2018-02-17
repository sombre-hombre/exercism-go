package raindrops

import (
	"bytes"
	"strconv"
)

var words = []struct {
	number int
	word   string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert a number to a string, the contents of which depend on the number's factors.
func Convert(n int) string {
	var buffer bytes.Buffer

	for _, item := range words {
		if n%item.number == 0 {
			buffer.WriteString(item.word)
		}
	}

	if buffer.Len() == 0 {
		buffer.WriteString(strconv.Itoa(n))
	}

	return buffer.String()
}
