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
