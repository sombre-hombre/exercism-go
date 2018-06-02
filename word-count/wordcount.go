package wordcount

import (
	"strings"
)

// Frequency is a map word to it's frequency in sentence
type Frequency map[string]int

// WordCount returns words frequency map
// BenchmarkWordCount-8   	  200000	      7776 ns/op	    3072 B/op	      28 allocs/op
func WordCount(phrase string) Frequency {
	f := make(Frequency)
	phrase = strings.ToLower(phrase)

	s := 0
	for {
		e := strings.IndexAny(phrase[s:], " ,")
		if e == -1 {
			f.add(trim(phrase[s:]))
			break
		}

		f.add(trim(phrase[s : s+e]))
		s += e + 1
	}

	return f
}

func trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return (r < 'a' || r > 'z') && (r < '0' || r > '9')
	})
}

func (f *Frequency) add(s string) {
	if len(s) > 0 {
		(*f)[s]++
	}
}
