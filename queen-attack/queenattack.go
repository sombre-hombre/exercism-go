package queenattack

import (
	"errors"
)

// CanQueenAttack given the position of two queens on a chess board,
// indicate whether or not they are positioned so that they can attack each other.
// BenchmarkCanQueenAttack-8   	 1000000	      1783 ns/op	      96 B/op	       6 allocs/op
func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("Same position")
	}

	white, ok := parse(w)
	if !ok {
		return false, errors.New("Invalid position for white queen")
	}

	black, ok := parse(b)
	if !ok {
		return false, errors.New("Invalid posotion for black queen")
	}

	return white.row == black.row || white.col == black.col || abs(white.row-black.row) == abs(white.col-black.col), nil
}

type position struct {
	col, row int
}

var rows = map[byte]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8}
var columns = map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8}

func parse(pos string) (p position, ok bool) {
	if len(pos) != 2 {
		return
	}

	c, ok := rows[pos[0]]
	if !ok {
		return
	}

	r, ok := columns[pos[1]]
	if !ok {
		return
	}

	return position{r, c}, true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
