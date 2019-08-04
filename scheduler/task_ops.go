package scheduler

import (
	"github.com/pkg/errors"
)

type taskManager struct {
	tasks           map[string]*Task
	disabledTasks   map[string]bool
	taskLastRunTime map[string]int64
}

func (tm *taskManager) initialize(tasks []*Task) {
	for _, v := range tasks {
		tm.tasks[v.Name] = v
	}
}

// GetTasks returns the currently loaded Tasks
func GetTasks() []*Task {
	tasks := make([]*Task, 0)

	for _, v := range tm.tasks {
		if b, ok := tm.disabledTasks[v.Name]; b && ok {
			continue
		}
		tasks = append(tasks, v)
	}

	return tasks
}

// GetTask returns a task by its name
func GetTask(name string) (*Task, error) {
	if t, ok := tm.tasks[name]; ok {
		return t, nil
	}
	return nil, errors.New("task not found")
}

// DisableTask removes a task from the running tasks slice and puts it in the disableTasks
func DisableTask(name string) error {
	var t *Task
	var ok bool

	if t, ok = tm.tasks[name]; !ok {
		return errors.New("No task by that name found")
	}

	tm.disabledTasks[t.Name] = true

	return nil
}

// GetDisabledTasks returns disabled tasks as a slice
// returns error if task is not found
func GetDisabledTasks() []*Task {
	tasks := make([]*Task, 0)
	for k, v := range tm.disabledTasks {
		if v {
			tasks = append(tasks, tm.tasks[k])
		}
	}
	return tasks
}

// EnableTask reenables a disabled
func EnableTask(name string) error {

	if _, ok := tm.tasks[name]; !ok {
		return errors.New("No task by that name found")
	}

	tm.disabledTasks[name] = false

	return nil
}
