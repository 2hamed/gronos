package scheduler

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

// Tasks is a type alias of []Task
type Tasks []*Task

// TaskMap is map to store references to Tasks
type TaskMap map[string]*Task

// Task is command which run by Schdule
type Task struct {
	Name     string   `yaml:"name"`
	Command  []string `yaml:"command"`
	Schedule Schedule `yaml:"schedule"`
}

var tasks Tasks
var taskLastRunTime = make(map[string]int64)
var taskMap = make(TaskMap)

// GetTasks returns the currently loaded Tasks
func GetTasks() Tasks {
	return tasks
}

// Execute executes the Task Command inside a separate goroutine
func (task Task) Execute() {
	go func() {
		taskLastRunTime[task.Name] = time.Now().Unix()
		log.Println("Running:", task.Command)
		cmd := exec.Command(task.Command[0], task.Command[1:]...)
		err := cmd.Run()
		if err != nil {
			log.Println(fmt.Errorf("Running command %s failed", task.Command))
			log.Println(err)
		}
	}()
}

// IsTime returns `true` if it's time to run this task
func (task Task) IsTime(anchor *time.Time) bool {

	result := false

	result = task.checkEvery(anchor)

	weekday, _ := task.Schedule.checkWeekday(anchor)
	result = result && weekday

	monthday, _ := task.Schedule.checkMonthdays(anchor)
	result = result && monthday

	at, _ := task.Schedule.checkAt(anchor)
	result = result && at

	months, _ := task.Schedule.checkMonths(anchor)
	result = result && months

	between, _ := task.Schedule.checkBetweens(anchor)
	result = result && between

	return result && !task.shouldSkip(anchor)
}

func (task Task) checkEvery(anchor *time.Time) bool {

	result := false

	every := task.Schedule.Every

	if lastTime, ok := taskLastRunTime[task.Name]; ok {
		if diff := time.Now().Unix() - lastTime; int(diff) > every {
			result = true
		}
	} else {
		result = true
	}

	return result
}

func (task Task) shouldSkip(anchor *time.Time) bool {

	schedule := task.Schedule

	except := schedule.Except

	if except == nil {
		return false
	}

	// `e` here inidicates whether the criteria was empty or not
	// and if it's empty it should not be considered a valid "true" for skipping the schedule

	if m, e := except.checkMonthdays(anchor); m && !e {
		return true
	}

	if w, e := except.checkWeekday(anchor); w && !e {
		return true
	}

	if a, e := except.checkAt(anchor); a && !e {
		return true
	}

	if m, e := except.checkMonths(anchor); m && !e {
		return true
	}

	if b, e := except.checkBetweens(anchor); b && !e {
		return true
	}

	return false
}
