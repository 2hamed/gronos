package main

import "time"

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
