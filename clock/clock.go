// Package clock implement a clock that handles times without dates
package clock

import (
	"fmt"
)

// Clock that handles times without dates
type Clock struct {
	hour, minute int
}

// New creates new Clock object
func New(hour, minute int) Clock {
	m := minute % 60
	if m < 0 {
		hour--
		m += 60
	}
	h := (minute/60 + hour%24) % 24
	if h < 0 {
		h += 24
	}

	return Clock{
		hour:   h,
		minute: m,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add minutes to clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract minutes from clock
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}
