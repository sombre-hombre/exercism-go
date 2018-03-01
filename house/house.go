package house

import (
	"fmt"
)

var lines = []string{
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

// Verse returns n-th verse of the song
func Verse(n int) string {
	if n == 0 || n > len(lines) {
		return ""
	}

	return fmt.Sprintf("This is %s", verse(n))
}

func verse(n int) string {
	if n == 0 {
		return ""
	}

	return lines[n-1] + verse(n-1)
}

// Song returns full song
func Song() string {
	verses := make([]string, len(lines))
	for i := 0; i < len(lines); i++ {
		verses[i] = Verse(i + 1)
	}

	return join("\n\n", verses) // or just strings.Join(verses, "\n\n")
}

func join(delimiter string, parts []string) string {
	if len(parts) == 0 {
		return ""
	}
	if len(parts) == 1 {
		return parts[0]
	}

	return parts[0] + delimiter + join(delimiter, parts[1:])
}
