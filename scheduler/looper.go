package scheduler

import (
	"time"

	log "github.com/sirupsen/logrus"
)

var tm = &taskManager{
	tasks:           make(map[string]*Task),
	disabledTasks:   make(map[string]bool),
	taskLastRunTime: make(map[string]int64),
}
var o *Options

// StartLooper starts the main looper for tasks
func StartLooper(initialConfigPath string, options ...SchedulerOption) {

	o = new(Options)

	for _, opt := range options {
		opt(o)
	}

	initStorage(o)

	var tasks tasks

	err := load(&tasks)

	if err != nil {
		tasks, err = LoadTasksFromDir(tm, initialConfigPath)
		if err != nil {
			panic(err)
		}
	}

	tm.initialize(tasks)

	ticker := time.Tick(1 * time.Second)

	go looper(tm.tasks, ticker)

	err = store(tasks)
	if err != nil {
		log.Errorf("Failed to persist tasks: %v", err)
	}
}

func looper(tasks Tasks, ticker <-chan time.Time) {
	for {
		t := <-ticker
		for _, task := range tasks {
			if task.IsTime(&t) {
				task.Execute()
			}
		}
	}
}
