package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix represents a matrix of numbers
type Matrix struct {
	rows, cols jaggedSlice
}

type jaggedSlice [][]int

// New creates new Matrix
func New(s string) (*Matrix, error) {
	if s == "" {
		return &Matrix{}, nil
	}

	m := Matrix{}
	rows := strings.Split(s, "\n")
	m.rows = make([][]int, len(rows))
	for i, row := range rows {
		cols := strings.Split(strings.TrimSpace(row), " ")
		m.rows[i] = make([]int, len(cols))

		if i > 0 && len(m.rows[i-1]) != len(m.rows[i]) {
			return nil, errors.New("Invalid data")
		}

		for j, cell := range cols {
			n, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			m.rows[i][j] = n
		}
	}

	m.cols = makeJaggedSlice(len(m.rows[0]), len(m.rows))
	for i := 0; i < len(m.rows[0]); i++ {
		for j := 0; j < len(m.rows); j++ {
			m.cols[i][j] = m.rows[j][i]
		}
	}

	return &m, nil
}

func makeJaggedSlice(r, c int) jaggedSlice {
	s := make([][]int, r)
	for i := 0; i < r; i++ {
		s[i] = make([]int, c)
	}

	return s
}

func (src jaggedSlice) Copy() jaggedSlice {
	dst := makeJaggedSlice(len(src), len(src[0]))
	for i := 0; i < len(src); i++ {
		copy(dst[i], src[i])
	}
	return dst
}

// Rows returns rows of matrix m
func (m Matrix) Rows() [][]int {
	return m.rows.Copy()
}

// Cols returns columns of matrix m
func (m Matrix) Cols() [][]int {
	return m.cols.Copy()
}

// Set sets a value to element of a matrix
func (m *Matrix) Set(i, j, val int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if len(m.cols) < i+1 || len(m.rows) < j+1 {
		return false
	}

	m.rows[i][j] = val
	m.cols[j][i] = val

	return true
}
