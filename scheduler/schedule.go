package scheduler

import (
	"strconv"
	"strings"
	"time"
)

// Schedule is struct containing when the task should be run
type Schedule struct {
	Months    []time.Month   `json:"months,omitempty"`
	Every     float64        `json:"every"`
	Weekdays  []time.Weekday `json:"weekdays,omitempty"`
	Monthdays []int          `json:"monthdays,omitempty"`
	At        []Hour         `json:"at,omitempty"`
	Betweens  []Between      `json:"betweens,omitempty"`
	Except    *Schedule      `json:"except,omitempty"`
}

// UnmarshalYAML custom YAML unmarshalling
func (s *Schedule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m map[interface{}]interface{}

	err := unmarshal(&m)

	if err != nil {
		return err
	}

	return s.populate(m)
}

func (s *Schedule) populate(m map[interface{}]interface{}) error {

	if v, ok := m["every"]; ok {
		var err error
		s.Every, err = parseEvery(v.(string))
		if err != nil {
			return err
		}
	}

	if v, ok := m["weekdays"]; ok {
		wds := v.([]interface{})
		weekdays := make([]string, len(wds))
		for i, v := range wds {
			weekdays[i] = v.(string)
		}
		s.Weekdays = parseWeekdays(weekdays)
	}

	if v, ok := m["monthdays"]; ok {
		mds := v.([]interface{})
		for _, v := range mds {
			s.Monthdays = append(s.Monthdays, v.(int))
		}
	}

	if v, ok := m["at"]; ok {
		ats := v.([]interface{})
		s.At = parseAt(ats)
	}

	exc := m["except"]
	if exc != nil {
		var except = Schedule{}
		_ = except.populate(exc.(map[interface{}]interface{}))
		s.Except = &except
	}

	months := m["months"]
	if months != nil {
		s.Months = make([]time.Month, 0)
		for _, v := range months.([]interface{}) {
			m, err := ParseMonth(v)
			if err == nil {
				s.Months = append(s.Months, m)
			}
		}
	}

	if b, ok := m["between"]; ok {
		betweens := b.([]interface{})
		s.Betweens = make([]Between, 0)
		if betweens != nil {
			for _, v := range betweens {
				b, err := parseBetween(v.(string))
				if err != nil {
					return err
				}
				s.Betweens = append(s.Betweens, b)
			}
		}
	}
	return nil
}

// At returns an array of time.Time struct in which only the Hour and Minute are important. The rest of the properties are arbitrary
func parseAt(ats []interface{}) []Hour {
	times := make([]Hour, len(ats))
	var s string
	var isString bool
	for i, v := range ats {
		if s, isString = v.(string); !isString {
			intg := v.(int)
			s = strconv.Itoa(intg)
		}
		hm := strings.Split(s, ":")
		hour, _ := strconv.Atoi(hm[0])
		var min int
		if len(hm) > 1 {
			min, _ = strconv.Atoi(hm[1])
		}
		times[i], _ = NewHour(hour, min)
	}

	return times
}

// Weekdays returns an array of time.Weekday
func parseWeekdays(wdays []string) []time.Weekday {
	weekdays := make([]time.Weekday, len(wdays))
	for i, wd := range wdays {
		switch wd {
		case "wed", "wedensday":
			weekdays[i] = time.Wednesday
			break
		case "sun", "sunday":
			weekdays[i] = time.Sunday
			break
		case "mon", "monday":
			weekdays[i] = time.Monday
			break
		case "tue", "tuesday":
			weekdays[i] = time.Tuesday
			break
		case "thu", "thursday":
			weekdays[i] = time.Thursday
			break
		case "sat", "saturday":
			weekdays[i] = time.Saturday
			break
		case "fri", "friday":
			weekdays[i] = time.Friday
			break
		}
	}
	return weekdays
}

// Every return the number of seconds at which the task should be run
func parseEvery(every string) (float64, error) {
	dur, err := time.ParseDuration(every)
	return dur.Seconds(), err
}

func (s Schedule) checkWeekday(anchor *time.Time) (bool, empty bool) {
	if len(s.Weekdays) == 0 {
		return true, true
	}
	return WeekDaySliceContains(s.Weekdays, anchor.Weekday()), false
}

func (s Schedule) checkMonthdays(anchor *time.Time) (bool, empty bool) {
	if len(s.Monthdays) == 0 {
		return true, true
	}
	return IntSliceContains(s.Monthdays, anchor.Day()), false
}

func (s Schedule) checkAt(anchor *time.Time) (bool, empty bool) {
	if len(s.At) == 0 {
		return true, true
	}
	return HourSliceContainsHoursMintues(s.At, *anchor), false
}

func (s Schedule) checkMonths(anchor *time.Time) (bool, empty bool) {
	if len(s.Months) == 0 {
		return true, true
	}
	return MonthSliceContains(s.Months, anchor.Month()), false
}

func (s Schedule) checkBetweens(anchor *time.Time) (bool, empty bool) {
	if len(s.Betweens) == 0 {
		return true, true
	}
	for _, b := range s.Betweens {
		if !b.IsInside(anchor) {
			return false, false
		}
	}
	return true, false
}
