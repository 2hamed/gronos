package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"gopkg.in/yaml.v2"
)

// Tasks is a type alias of []Task
type Tasks []*Task

// TaskMap is map to store references to Tasks
type TaskMap map[string]*Task

// Task is command which run by Schdule
type Task struct {
	Name     string   `yaml:"name"`
	Command  string   `yaml:"command"`
	Schedule Schedule `yaml:"schedule"`
}

var taskLastRunTime = make(map[string]int64)
var taskMap = make(TaskMap)

// Execute executes the Task Command inside a separate goroutine
func (task Task) Execute() {
	go func() {
		taskLastRunTime[task.Name] = time.Now().Unix()
		err := syscall.Exec(task.Command, nil, nil)
		if err != nil {
			log.Println(fmt.Errorf("Running command %s failed: %v", task.Command, err))
		}
	}()
}

// IsTime returns `true` if it's time to run this task
func (task Task) IsTime() bool {

	result := false

	schedule := task.Schedule
	every, err := schedule.Every()

	now := time.Now()

	if err == nil {
		if lastTime, ok := taskLastRunTime[task.Name]; ok {
			if diff := time.Now().Unix() - lastTime; diff > every {
				result = true
			}
		} else {
			result = true
		}
	}

	if WeekDaySliceContains(schedule.Weekdays(), now.Weekday()) {
		result = result && true
	}

	if IntSliceContains(schedule.Monthdays(), now.Day()) {
		result = result && true
	}

	if TimeSliceContainsHoursMintues(schedule.At(), now) {
		result = result && true
	}
	// TODO: check for other criteria

	return result && !task.shouldSkip()
}

func (task Task) shouldSkip() bool {

	schedule := task.Schedule

	except := schedule.Except()

	now := time.Now()

	if except == nil {
		return false
	}

	if IntSliceContains(except.monthdays, int(now.Month())) {
		return true
	}

	if WeekDaySliceContains(except.Weekdays(), now.Weekday()) {
		return true
	}

	for _, t := range except.At() {
		if t.Hour() == now.Hour() && t.Minute() == now.Minute() {
			return true
		}
	}

	return false
}

// LoadTasksFromFile reads an etire YAML file and outputs the corresponding Tasks struct
func LoadTasksFromFile(filePath string) (Tasks, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

// LoadTasksFromDir scans a directory and loads every YAML file into corresponding Tasks struct
func LoadTasksFromDir(dirPath string) (Tasks, error) {

	dirPath, _ = filepath.Abs(dirPath)

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var tasks Tasks = make(Tasks, 0)

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yaml") {
			ts, err := LoadTasksFromFile(dirPath + "/" + f.Name())
			if err != nil {
				panic(err)
			}

			for _, task := range ts {
				if _, ok := taskMap[task.Name]; ok {
					return nil, errors.New("duplicate task name: " + task.Name)
				}
				tasks = append(tasks, task)
				taskMap[task.Name] = task
			}
		}
	}

	return tasks, nil
}
