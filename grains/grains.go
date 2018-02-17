package grains

import (
	"errors"
)

// Square calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
func Square(cell int) (uint64, error) {
	if cell > 64 || cell < 1 {
		return 0, errors.New("Invalid cell")
	}

	return 1 << uint(cell-1), nil
}

// Total calculates total number of grains of wheat on a chessboard given that the number on each square doubles.
func Total() uint64 {
	var result uint64
	for i := uint(0); i < 64; i++ {
		result += 1 << i
	}

	return result
}
