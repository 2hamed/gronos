package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Schedule is struct containing when the task should be run
type Schedule struct {
	name      string   `yaml:"name"`
	every     string   `yaml:"every"`
	weekdays  []string `yaml:"weekdays"`
	monthdays []int    `yaml:"monthdays"`
	at        []string `yaml:"at"`

	except *Schedule `yaml:"except"`
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
		s.every = v.(string)
	}

	if v, ok := m["weekdays"]; ok {
		wds := v.([]interface{})
		for _, v := range wds {
			s.weekdays = append(s.weekdays, v.(string))
		}
	}

	if v, ok := m["monthdays"]; ok {
		mds := v.([]interface{})
		for _, v := range mds {
			s.monthdays = append(s.monthdays, int8(v.(int)))
		}
	}

	if v, ok := m["at"]; ok {
		ats := v.([]interface{})
		for _, v := range ats {
			if str, ok := v.(string); ok {
				s.at = append(s.at, str)
			} else if intg, ok := v.(int); ok {
				s.at = append(s.at, strconv.Itoa(intg))
			} else {
				return fmt.Errorf("invalid value for `at`: %v", v)

			}
		}
	}

	exc := m["except"]
	if exc != nil {
		var except = Schedule{}
		except.populate(exc.(map[interface{}]interface{}))
		s.except = &except
	}

	return nil
}

// At returns an array of time.Time struct in which only the Hour and Minute are important. The rest of the properties are arbitrary
func (s Schedule) At() []time.Time {
	times := make([]time.Time, len(s.at))
	for i, v := range s.at {
		hm := strings.Split(v, ":")
		hour, _ := strconv.Atoi(hm[0])
		var min int
		if len(hm) > 1 {
			min, _ = strconv.Atoi(hm[1])
		}
		times[i] = time.Date(1, 1, 1, hour, min, 0, 0, time.Local)
	}

	return times
}

// Weekdays returns an array of time.Weekday
func (s Schedule) Weekdays() []time.Weekday {
	weekdays := make([]time.Weekday, len(s.weekdays))
	for i, wd := range s.weekdays {
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

// Except returns the Schedule struct denoting when not to run this task
func (s Schedule) Except() *Schedule {
	return s.except
}

// Every return the number of seconds at which the task should be run
func (s Schedule) Every() (int64, error) {
	hm := strings.Split(s.every, ":")
	if len(hm) == 0 {
		return 0, errors.New("no `every` set")
	}
	hour, _ := strconv.Atoi(hm[0])
	var min int
	if len(hm) > 1 {
		min, _ = strconv.Atoi(hm[1])
	}

	return int64(hour*3600 + min*60), nil
}

// IsTime checks whether the Schedule is set to run on `t` or not
func (s Schedule) IsTime(taskName string, t time.Time) bool {

	return true
}
