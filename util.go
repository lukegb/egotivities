package eactivities

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	// The point at which the year transitions
	cutoffMonth = time.August
)

var (
	eActivitiesLocation = mustLocation(time.LoadLocation("Europe/London"))
)

func mustLocation(l *time.Location, err error) *time.Location {
	if err != nil {
		panic(err)
	}
	return l
}

// DateToYear returns the eActivities year string containing the given date.
func DateToYear(t time.Time) string {
	t = t.In(eActivitiesLocation)

	startYear := t.Year()
	if t.Month() < cutoffMonth {
		// This is the end of the year
		startYear--
	}
	startYear = (startYear % 100)
	return fmt.Sprintf("%d-%d", startYear, startYear+1)
}

// CurrentYear returns the current eActivities year string.
func CurrentYear() string {
	return DateToYear(time.Now())
}

// Time represents a Time returned by eActivities' API.
type Time time.Time

// UnmarshalJSON handles deserializing eActivities' special date/time format.
func (t *Time) UnmarshalJSON(date []byte) error {
	var x string
	if err := json.Unmarshal(date, &x); err != nil {
		return err
	}
	parseTmpl := "2006-01-02"
	if strings.ContainsRune(x, ' ') {
		parseTmpl += " 15:04:05"
	}

	parsed, err := time.ParseInLocation(parseTmpl, x, eActivitiesLocation)
	if err != nil {
		return err
	}
	*t = Time(parsed)
	return nil
}
