package scheduler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
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

// HourSliceContainsHoursMintues checks an instance of time inside a slice only for the hour and minute
func HourSliceContainsHoursMintues(haystack []Hour, n time.Time) bool {
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

// ParseMonth parses a string or int value and converts it to a time.Month
func ParseMonth(v interface{}) (time.Month, error) {
	var m time.Month
	var err error

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

// LoadTasksFromFile reads an etire YAML file and outputs the corresponding Tasks struct
func LoadTasksFromFile(filePath string) ([]*Task, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

// LoadTasksFromDir scans a directory and loads every YAML file into corresponding Tasks struct
func LoadTasksFromDir(tm *taskManager, dirPath string) ([]*Task, error) {

	dirPath, _ = filepath.Abs(dirPath)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var tasks = make([]*Task, 0)

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yaml") {
			ts, err := LoadTasksFromFile(dirPath + "/" + f.Name())
			if err != nil {
				panic(err)
			}

			for _, task := range ts {
				if _, ok := tm.tasks[task.Name]; ok {
					return nil, errors.New("duplicate task name: " + task.Name)
				}
				tasks = append(tasks, task)
				tm.tasks[task.Name] = task
			}
		}
	}

	return tasks, nil
}
