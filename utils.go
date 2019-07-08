package main

import (
	"fmt"
	"time"
)

// Generics would have been nice

// IntSliceContains checks a value `n` exists in slice `haystack`
func IntSliceContains(haystack []int, n int) bool {

	if haystack == nil {
		return false
	}

	for _, v := range haystack {
		if v == n {
			return true
		}
	}
	return false
}

// WeekDaySliceContains checks a value `n` exists in slice `haystack`
func WeekDaySliceContains(haystack []time.Weekday, n time.Weekday) bool {

	if haystack == nil {
		return false
	}

	for _, v := range haystack {
		if v == n {
			return true
		}
	}
	return false
}

// TimeSliceContainsHoursMintues checks an instance of time inside a slice only for the hour and minute
func TimeSliceContainsHoursMintues(haystack []time.Time, n time.Time) bool {
	if haystack == nil {
		return false
	}

	for _, t := range haystack {
		if t.Hour() == n.Hour() && t.Minute() == n.Minute() {
			return true
		}
	}

	return false
}

// MonthSliceContains hecks a value `n` exists in slice `haystack`
func MonthSliceContains(haystack []time.Month, n time.Month) bool {
	if haystack == nil {
		return false
	}

	for _, t := range haystack {
		if t == n {
			return true
		}
	}

	return false
}

func ParseMonth(v interface{}) (time.Month, error) {
	var m time.Month
	var err error = nil

	switch v {
	case "jan", "january", 1:
		m = time.January
	case "feb", "february", 2:
		m = time.February
	case "mar", "march", 3:
		m = time.March
	case "apr", "april", 4:
		m = time.April
	case "may", 5:
		m = time.May
	case "jun", "june", 6:
		m = time.June
	case "jul", "july", 7:
		m = time.July
	case "aug", "august", 8:
		m = time.August
	case "sep", "september", 9:
		m = time.September
	case "oct", "october", 10:
		m = time.October
	case "nov", "november", 11:
		m = time.November
	case "dec", "december", 12:
		m = time.December
	default:
		err = fmt.Errorf("invalid input %v", v)
	}

	return m, err
}
