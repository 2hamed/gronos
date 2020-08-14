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

// StartLooper starts the main looper for tasks
func StartLooper(configPath string) {
	var tasks []*Task

	err := load(&tasks)
	if err != nil {
		tasks, err = LoadTasksFromDir(tm, configPath)
		if err != nil {
			panic(err)
		}
	}

	tm.initialize(tasks)

	ticker := time.Tick(1 * time.Second)

	go looper(tm.tasks, ticker)

	err = persist(tasks)
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
