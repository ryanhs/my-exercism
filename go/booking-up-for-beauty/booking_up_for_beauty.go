package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dateLayout := "1/2/2006 15:04:05"
	t, _ := time.Parse(dateLayout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateLayout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(dateLayout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateLayout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(dateLayout, date)
	hour := t.Hour()
	// fmt.Printf("INFO: IsAfternoonAppointment: %s > h %d\n", date, hour)
	return 12 <= hour && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dateLayout := "1/2/2006 15:04:05"
	t, _ := time.Parse(dateLayout, date)
	newDateFormat := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", newDateFormat)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2020-09-15 00:00:00 +0000 UTC")
	return time.Date(
		time.Now().Year(), // update to current year
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		t.Location(),
	)
}
