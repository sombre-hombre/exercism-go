package foodchain

import (
	"fmt"
	"strings"
)

var lines = []struct {
	a, b string
}{
	{"fly.", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider.\nIt wriggled and jiggled and tickled inside her.", "She swallowed the spider to catch the fly."},
	{"bird.\nHow absurd to swallow a bird!", "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her."},
	{"cat.\nImagine that, to swallow a cat!", "She swallowed the cat to catch the bird."},
	{"dog.\nWhat a hog, to swallow a dog!", "She swallowed the dog to catch the cat."},
	{"goat.\nJust opened her throat and swallowed a goat!", "She swallowed the goat to catch the dog."},
	{"cow.\nI don't know how she swallowed a cow!", "She swallowed the cow to catch the goat."},
	{"horse.", "She's dead, of course!"},
}

// Verse returns n-th verse of the song
func Verse(n int) string {
	if n < 1 || n > len(lines) {
		return ""
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "I know an old lady who swallowed a %s\n", lines[n-1].a)
	verse(n, &sb)

	return sb.String()
}

func verse(n int, sb *strings.Builder) {
	sb.WriteString(lines[n-1].b)

	if n == len(lines) {
		return
	}

	if n > 1 {
		sb.WriteString("\n")
		verse(n-1, sb)
	}
}

// Verses returns verses from n-th to m-th
func Verses(n, m int) string {
	if n < 1 || n >= len(lines) {
		return ""
	}

	var sb strings.Builder
	for i := n; i <= m; i++ {
		sb.WriteString(Verse(i))
		if i < m {
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}

// Song returns full song
func Song() string {
	return Verses(1, len(lines))
}

// goos: windows
// goarch: amd64
// BenchmarkSong-8   	  200000	      6839 ns/op	   13461 B/op	      47 allocs/op
// PASS
