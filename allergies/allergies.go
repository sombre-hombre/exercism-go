package allergies

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var allergies = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

var scores = map[uint]string{
	eggs:         "eggs",
	peanuts:      "peanuts",
	shellfish:    "shellfish",
	strawberries: "strawberries",
	tomatoes:     "tomatoes",
	chocolate:    "chocolate",
	pollen:       "pollen",
	cats:         "cats",
}

// BenchmarkAllergies-8   	 5000000	       283 ns/op	      64 B/op	       3 allocs/op
func Allergies(score uint) []string {
	r := make([]string, 0)
	for i := 0; i < len(allergies); i++ {
		var s uint = 1 << uint(i)
		if score&s != 0 {
			r = append(r, scores[s])
		}
	}

	return r
}

// BenchmarkAllergicTo-8   	20000000	        78.1 ns/op	       0 B/op	       0 allocs/op
func AllergicTo(i uint, s string) bool {
	if score, ok := allergies[s]; ok {
		return i&score != 0
	}
	return false
}
