package diamond

import (
	"errors"
	"strings"
)

// Gen generates the Diamond kata
func Gen(letter byte) (string, error) {
	if letter < 'A' || letter > 'Z' {
		return "", errors.New("Illegal character")
	}

	delta := int(letter - 'A')
	size := int(delta*2 + 1)
	middle := size / 2
	var m [][]byte
	var sb strings.Builder

	row := []byte(strings.Repeat(" ", size))

	m = make([][]byte, middle+1)

	for i := 0; i <= middle; i++ {
		m[i] = make([]byte, size)
		copy(m[i], row)
		m[i][middle+i] = byte('A' + i)
		m[i][middle-i] = byte('A' + i)

		sb.Write(m[i])
		sb.WriteByte('\n')
	}

	for i := middle + 1; i < size; i++ {
		sb.Write(m[size-1-i])
		sb.WriteByte('\n')
	}

	return sb.String(), nil
}
