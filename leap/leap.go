// Package leap should have a package comment that summarizes what it's about.
package leap

// IsLeapYear report if it is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
