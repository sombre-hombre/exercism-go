// Package twofer implements "Two Fer in Go" excercise
package twofer

import "fmt"

// ShareWith creates a sentence of the form "One for X, one for me."
func ShareWith(who string) string {
	if who == "" {
		who = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", who)
}
