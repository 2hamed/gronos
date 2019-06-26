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

func ParseMonth(v interface{}) (time.Month, error) {
	switch v {
	case "jan", "january", 1:
		return time.January, nil
	case "feb", "february", 2:
		return time.February, nil
	case "mar", "march", 3:
		return time.March, nil
	case "apr", "april", 4:
		return time.April, nil
	case "may", 5:
		return time.May, nil
	case "jun", "june", 6:
		return time.June, nil
	case "jul", "july", 7:
		return time.July, nil
	case "aug", "august", 8:
		return time.August, nil
	case "sep", "september", 9:
		return time.September, nil
	case "oct", "october", 10:
		return time.October, nil
	case "nov", "november", 11:
		return time.November, nil
	case "dec", "december", 12:
		return time.December, nil
	default:
		return time.January, fmt.Errorf("invalid input %v", v)
	}
}
