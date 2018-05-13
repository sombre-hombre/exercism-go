package scale

import (
	"strings"
)

var sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var notesToScales = map[string][]string{
	"a":  sharpScale,
	"C":  sharpScale,
	"G":  sharpScale,
	"D":  sharpScale,
	"A":  sharpScale,
	"E":  sharpScale,
	"B":  sharpScale,
	"F#": sharpScale,
	"e":  sharpScale,
	"b":  sharpScale,
	"f#": sharpScale,
	"c#": sharpScale,
	"g#": sharpScale,
	"d#": sharpScale,
	"F":  flatScale,
	"Bb": flatScale,
	"Eb": flatScale,
	"Ab": flatScale,
	"Db": flatScale,
	"Gb": flatScale,
	"d":  flatScale,
	"g":  flatScale,
	"с":  flatScale,
	"f":  flatScale,
	"bb": flatScale,
	"eb": flatScale,
}

// Scale generates the musical scale starting with the tonic and following the specified interval pattern.
// BenchmarkScale-8   	  200000	      6635 ns/op	    3152 B/op	      34 allocs/op
func Scale(tonic string, interval string) []string {
	intervals := parseIntervals(interval)
	scale := notesToScales[tonic]

	result := make([]string, len(intervals))
	i := indexOf(scale, tonic)
	for j, step := range intervals {
		result[j] = scale[i%len(scale)]
		i += step
	}

	return result
}

func parseIntervals(interval string) []int {
	if len(interval) == 0 {
		return []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	}

	r := make([]int, len(interval))
	for i, s := range interval {
		switch s {
		case 'm':
			r[i] = 1
		case 'M':
			r[i] = 2
		case 'A':
			r[i] = 3
		}
	}

	return r
}

func indexOf(s []string, el string) int {
	for i, e := range s {
		if strings.EqualFold(e, el) {
			return i
		}
	}

	return -1
}
