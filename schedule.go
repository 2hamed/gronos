package main

import (
	"fmt"
	"strconv"
)

// Schedule is struct containing when the task should be run
type Schedule struct {
	name      string   `yaml:"name"`
	every     string   `yaml:"every"`
	weekdays  []string `yaml:"weekdays"`
	monthdays []int8   `yaml:"monthdays"`
	at        []string `yaml:"at"`

	except *Schedule `yaml:"except"`
}

func (s *Schedule) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var m map[interface{}]interface{}

	err := unmarshal(&m)

	if err != nil {
		return err
	}

	return s.populate(m)
}

func (s *Schedule) populate(m map[interface{}]interface{}) error {
	if v, ok := m["name"]; ok {
		s.name = v.(string)
	}

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

// func (s Schedule) At() []time.Time {
// 	times := make([]time.Time, len(s.at))
// 	for i, v := range s.at {
// 		fmt.Println(v)
// 		t, err := time.Parse(time.RFC822, v)
// 		if err == nil {
// 			times[i] = t
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}

// 	return times
// }
func (s Schedule) Except() *Schedule {
	return s.except
}
