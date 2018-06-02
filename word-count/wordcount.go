package wordcount

import (
	"strings"
)

// Frequency is a map word to it's frequency in sentence
type Frequency map[string]int

// WordCount returns words frequency map
// BenchmarkWordCount-8   	  200000	      9139 ns/op	    3920 B/op	      42 allocs/op
func WordCount(phrase string) Frequency {
	f := make(map[string]int)

	words := strings.Split(strings.ToLower(phrase), " ")
	if len(words) == 1 {
		words = strings.Split(words[0], ",")
	}
	for _, w := range words {
		w = strings.TrimFunc(w, func(r rune) bool {
			return (r < 'a' || r > 'z') && (r < '0' || r > '9')
		})

		if w != "" {
			f[w]++
		}
	}

	return f
}
