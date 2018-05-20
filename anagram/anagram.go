package anagram

import (
	"strings"
)

// Detect finds a list of possible anagrams of the subject from the candidates.
// BenchmarkDetectAnagrams-8   	   50000	     31640 ns/op	    8032 B/op	     167 allocs/op
func Detect(subject string, candidates []string) []string {
	var letters = map[rune]int{}
	for _, r := range strings.ToLower(subject) {
		letters[r]++
	}

	result := make([]string, 0)
	for _, c := range candidates {
		if strings.EqualFold(subject, c) || len(c) != len(subject) {
			continue
		}

		if isAnagram(copy(letters), c) {
			result = append(result, c)
		}
	}

	return result
}

func isAnagram(letters map[rune]int, candidate string) bool {
	for _, r := range strings.ToLower(candidate) {
		if n, ok := letters[r]; !ok || n == 0 {
			return false
		}
		letters[r]--
	}

	for _, v := range letters {
		if v > 0 {
			return false
		}
	}

	return true
}

func copy(m map[rune]int) map[rune]int {
	c := make(map[rune]int, len(m))
	for key, value := range m {
		c[key] = value
	}

	return c
}
