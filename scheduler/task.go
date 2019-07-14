package scheduler

import (
	"fmt"
	"log"
	"syscall"
	"time"
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
func (task Task) IsTime(anchor *time.Time) bool {

	result := false

	result = task.checkEvery(anchor)

	result = result && task.Schedule.checkWeekday(anchor)

	result = result && task.Schedule.checkMonthdays(anchor)

	result = result && task.Schedule.checkAt(anchor)

	result = result && task.Schedule.checkMonths(anchor)

	return result && !task.shouldSkip(anchor)
}

func (task Task) checkEvery(anchor *time.Time) bool {

	result := false

	every, err := task.Schedule.Every()

	if err == nil {
		if lastTime, ok := taskLastRunTime[task.Name]; ok {
			if diff := time.Now().Unix() - lastTime; diff > every {
				result = true
			}
		} else {
			result = true
		}
	}

	return result
}

func (task Task) shouldSkip(anchor *time.Time) bool {

	schedule := task.Schedule

	except := schedule.Except()

	if except == nil {
		return false
	}

	if except.checkMonthdays(anchor) {
		return true
	}

	if except.checkWeekday(anchor) {
		return true
	}

	if except.checkAt(anchor) {
		return true
	}

	if except.checkMonths(anchor) {
		return true
	}

	return false
}
