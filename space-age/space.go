// Package space implements "Space Age in Go" exercise.
package space

// Planet — solar system planet
type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   1,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const secondsInYear = 31557600

// Age calculates how old someone is in terms of a given planet's solar years.
func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsInYear / orbitalPeriods[planet]
}
