package anagram

import (
	"strings"
)

// Detect finds a list of possible anagrams of the subject from the candidates.
// BenchmarkDetectAnagrams-8   	  100000	     14094 ns/op	    3040 B/op	      75 allocs/op
func Detect(subject string, candidates []string) []string {
	var letters = map[rune]int{}
	var subjectSumm int
	for _, r := range strings.ToLower(subject) {
		subjectSumm += int(r)
		letters[r]++
	}

	result := make([]string, 0)
	for _, c := range candidates {
		if strings.EqualFold(subject, c) || len(c) != len(subject) {
			continue
		}

		var sum int
		for _, r := range strings.ToLower(c) {
			sum += int(r)
		}
		if sum != subjectSumm {
			continue
		}

		if isAnagram(letters, c) {
			result = append(result, c)
		}
	}

	return result
}

func isAnagram(letters map[rune]int, candidate string) bool {
	letters = copy(letters)
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
