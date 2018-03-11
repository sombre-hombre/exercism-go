package etl

import (
	"strings"
)

// Transform converts scrabble scores from a legacy format
func Transform(scoreTable map[int][]string) map[string]int {
	result := make(map[string]int)
	for score, letters := range scoreTable {
		for _, letter := range letters {
			result[strings.ToLower(letter)] = score
		}
	}

	return result
}
