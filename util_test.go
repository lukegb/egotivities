package egotivities

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateToYear(t *testing.T) {
	testCases := []struct {
		name string
		date string
		want string
	}{
		{
			name: "Exactly changeover",
			date: "1 August 2016 00:00:00 +0100",
			want: "16-17",
		},
		{
			name: "Changeover... in France",
			date: "1 August 2016 00:00:00 +0200",
			want: "15-16",
		},
		{
			name: "End of the year",
			date: "31 December 2016 00:00:00 +0000",
			want: "16-17",
		},
		{
			name: "Start of the year",
			date: "1 January 2017 00:00:00 +0000",
			want: "16-17",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			date, err := time.Parse("2 January 2006 15:04:05 -0700", testCase.date)
			if err != nil {
				t.Fatalf("time.Parse(%q): %v", testCase.date, err)
			}
			got := DateToYear(date)
			if got != testCase.want {
				t.Fatalf("DateToYear(%s) = %q; want %q", date.String(), got, testCase.want)
			}
		})
	}
}

func TestTimeParsing(t *testing.T) {
	testCases := []struct {
		inp  string
		want time.Time
	}{
		{
			inp:  "\"2006-01-02\"",
			want: time.Date(2006, 1, 2, 0, 0, 0, 0, eActivitiesLocation),
		},
		{
			inp:  "\"2006-01-02 15:04:05\"",
			want: time.Date(2006, 1, 2, 15, 4, 5, 0, eActivitiesLocation),
		},
	}
	for _, testCase := range testCases {
		var got Time
		if err := json.Unmarshal([]byte(testCase.inp), &got); err != nil {
			t.Errorf("json.Unmarshal(%q, &got): %v", testCase.inp, err)
			continue
		}
		gotTime := time.Time(got)
		if !testCase.want.Equal(gotTime) {
			t.Errorf("json.Unmarshal(%q) = %s; want %s", testCase.inp, got, testCase.want)
		}
	}
}
