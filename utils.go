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

func ParseMonth(v interface{}) time.Month {
	switch v {
	case "jan", "january", 1:
		return time.January
	case "feb", "february", 2:
		return time.February
	case "mar", "march", 3:
		return time.March
	case "apr", "april", 4:
		return time.April
	case "may", 5:
		return time.May
	case "jun", "june", 6:
		return time.June
	case "jul", "july", 7:
		return time.July
	case "aug", "august", 8:
		return time.August
	case "sep", "september", 9:
		return time.September
	case "oct", "october", 10:
		return time.October
	case "nov", "november", 11:
		return time.November
	case "dec", "december", 12:
		return time.December
	default:
		panic("invalid month")
	}
}
