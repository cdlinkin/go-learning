package stage5generics

import "time"

func NextWorkday(t time.Time) time.Time {
	next := t.AddDate(0, 0, 1)

	switch next.Weekday() {
	case time.Saturday:
		next = next.AddDate(0, 0, 2)
	case time.Sunday:
		next = next.AddDate(0, 0, 1)
	}

	return next
}
