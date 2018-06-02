package wordcount

import (
	"strings"
)

// Frequency is a map word to it's frequency in sentence
type Frequency map[string]int

// WordCount returns words frequency map
// BenchmarkWordCount-8   	  200000	      9139 ns/op	    3920 B/op	      42 allocs/op
func WordCount1(phrase string) Frequency {
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

// WordCount returns words frequency map
// BenchmarkWordCount-8   	  200000	      9956 ns/op	    4784 B/op	      63 allocs/op
func WordCount(phrase string) Frequency {
	f := make(map[string]int)
	for _, word := range split(phrase) {
		f[word]++
	}

	return f
}

func split(s string) (r []string) {
	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		if !isWordChar(s[i]) {
			continue
		}

		if word := getWord(s[i:]); len(word) > 0 {
			r = append(r, word)
			i += len(word)
		}
	}

	return r
}

func isWordChar(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func getWord(s string) string {
	var i int
	for ; i < len(s); i++ {
		if isWordChar(s[i]) {
			continue
		}

		if s[i] == '\'' && i+1 < len(s) && isWordChar(s[i+1]) {
			continue
		}

		break
	}

	return s[:i]
}
