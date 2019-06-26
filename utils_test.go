package main

import (
	"testing"
	"time"
)

func TestIntSliceContains(t *testing.T) {
	haystack := []int{1, 3, 4, 6, 712, 45, -78}

	if !IntSliceContains(haystack, 4) {
		t.Error("slice contains 4, but it says it can't find it")
	}

	if IntSliceContains(haystack, 5) {
		t.Error("slice doesn't contain 5, but it says it does")
	}

	if IntSliceContains(nil, 5) {
		t.Error("nil slice should always return false")
	}
}

func TestWeekdaySliceContains(t *testing.T) {
	haystack := []time.Weekday{time.Thursday, time.Wednesday}

	if WeekDaySliceContains(haystack, time.Saturday) {
		t.Error("haystack does not contain time.Saturday")
	}

	if !WeekDaySliceContains(haystack, time.Thursday) {
		t.Error("haystack does contain time.Thursday")
	}

	if WeekDaySliceContains(nil, time.Saturday) {
		t.Error("nil slice should always return false")
	}
}

func TestTimeSliceContainsHoursMintues(t *testing.T) {
	haystack := []time.Time{
		time.Date(0, 0, 0, 13, 13, 0, 0, time.Local),
		time.Date(0, 0, 0, 3, 0, 0, 0, time.Local),
		time.Date(2010, 1, 1, 19, 1, 0, 0, time.Local),
		time.Date(0, 0, 0, 5, 59, 0, 0, time.Local),
	}

	if !TimeSliceContainsHoursMintues(haystack, time.Date(4, 4, 4, 13, 13, 10, 0, time.Local)) {
		t.Error("this should be true")
	}

	if TimeSliceContainsHoursMintues(haystack, time.Date(2010, 1, 1, 19, 0, 0, 0, time.Local)) {
		t.Error("this should be false")
	}

	if TimeSliceContainsHoursMintues(nil, time.Date(4, 5, 6, 73, 4, 6, 4, time.Local)) {
		t.Error("nil slice should always return false")
	}
}

func TestParseMonth(t *testing.T) {
	m, err := ParseMonth(1)
	if err != nil {
		t.Error(err)
	}

	m, err = ParseMonth("jun")
	if err != nil {
		t.Error(err)
	}

	m, err = ParseMonth("january")
	if err != nil {
		t.Error(err)
	}

	m, err = ParseMonth("")
	if err == nil {
		t.Error("should return error")
	}
}
