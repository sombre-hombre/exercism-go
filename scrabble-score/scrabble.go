package scrabble

import (
	"strings"
)

// Score compute the scrabble score
func Score(text string) int {
	result := 0
	// for each input symbol
	for _, r := range strings.ToUpper(text) {
		// for each score item
		for _, s := range scores {
			if s.isMatched(r) {
				result += s.score
			}
		}
	}

	return result
}

type scoreItem struct {
	runes []rune
	score int
}

var scores = []scoreItem{
	{[]rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}, 1},
	{[]rune{'D', 'G'}, 2},
	{[]rune{'B', 'C', 'M', 'P'}, 3},
	{[]rune{'F', 'H', 'V', 'W', 'Y'}, 4},
	{[]rune{'K'}, 5},
	{[]rune{'J', 'X'}, 8},
	{[]rune{'Q', 'Z'}, 10},
}

func (item scoreItem) isMatched(symbol rune) bool {
	for _, r := range item.runes {
		if symbol == r {
			return true
		}
	}
	return false
}
