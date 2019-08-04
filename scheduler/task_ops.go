package scheduler

import "github.com/pkg/errors"

var tasks Tasks
var taskLastRunTime = make(map[string]int64)
var taskMap = make(TaskMap)
var disabledTasks = make(map[string]*Task)

// GetTasks returns the currently loaded Tasks
func GetTasks() Tasks {
	return tasks
}

// GetTask returns a task by its name
func GetTask(name string) (*Task, error) {
	if t, ok := taskMap[name]; ok {
		return t, nil
	}
	return nil, errors.New("task not found")
}

// DisableTask removes a task from the running tasks slice and puts it in the disableTasks
func DisableTask(name string) error {
	var t *Task
	var ok bool

	if t, ok = taskMap[name]; !ok {
		return errors.New("No task by that name found")
	}

	disabledTasks[name] = t

	delete(taskMap, name)
	index := findTaskIndex(name)
	deleteTaskFromList(index)
	return nil
}

// GetDisabledTasks returns disabled tasks as a slice
func GetDisabledTasks() Tasks {
	tasks := make(Tasks, 0)
	for _, v := range disabledTasks {
		tasks = append(tasks, v)
	}
	return tasks
}

func findTaskIndex(name string) int {
	var i = -1
	for i = 0; i < len(tasks); i++ {
		if tasks[i].Name == name {
			break
		}
	}
	return i
}

func deleteTaskFromList(index int) {
	switch len(tasks) {
	case 0:
		break
	case 1:
		tasks = make(Tasks, 0)
		break
	default:
		tasks[index] = tasks[len(tasks)-1]
		tasks = tasks[:len(tasks)-1]
		break
	}
}
