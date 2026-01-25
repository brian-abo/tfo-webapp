package components

import "time"

// CurrentYear returns the current year as a string.
func CurrentYear() string {
	return time.Now().Format("2006")
}
