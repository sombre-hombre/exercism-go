package triangle

import (
	"math"
)

// Kind of a triangle
type Kind int

const (
	// NaT — not a triangle
	NaT Kind = iota
	// Equ — equilateral triangle
	Equ
	// Iso — isosceles triangle
	Iso
	// Sca — scalene triangle
	Sca
	// Deg — degenerate triangle
	Deg
)

// KindFromSides determines the kind of a triangle.
func KindFromSides(a, b, c float64) Kind {
	for _, s := range []float64{a, b, c} {
		if math.IsNaN(s) || math.IsInf(s, 1) || math.IsInf(s, -1) {
			return NaT
		}
	}

	// All sides have to be of length > 0.
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// The sum of the lengths of any two sides must be greater than or equal
	// to the length of the third side.
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	switch {
	case a == b && b == c:
		return Equ
	case math.Max(c, math.Max(a, b))*2 == a+b+c:
		return Deg
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
