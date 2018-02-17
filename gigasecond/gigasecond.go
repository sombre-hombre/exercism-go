// Package gigasecond implements Gigasecond exercise.
package gigasecond

import "time"

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds..
func AddGigasecond(t time.Time) time.Time {
	result := t.Add(time.Duration(1e9) * time.Second)

	return result
}
