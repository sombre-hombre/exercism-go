package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var letters = map[rune]int{}
	for _, r := range strings.ToLower(subject) {
		letters[r]++
	}

	result := make([]string, 0)
	for _, c := range candidates {
		// word is not own anagram
		if strings.EqualFold(subject, c) {
			break
		}

		isAnagram := true
		copy := copy(letters)
		for _, r := range strings.ToLower(c) {
			if n, ok := copy[r]; !ok || n == 0 {
				isAnagram = false
				break
			}
			copy[r]--
		}

		for _, v := range copy {
			if v > 0 {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			result = append(result, c)
		}
	}

	return result
}

func copy(m map[rune]int) map[rune]int {
	c := make(map[rune]int, len(m))
	for key, value := range m {
		c[key] = value
	}

	return c
}
