package main

import "time"

// InTimeSpan returns whether the check time falls within the start and end dates.
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
