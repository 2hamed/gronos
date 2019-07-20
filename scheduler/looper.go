package scheduler

import (
	"os"
	"time"
)

func init() {
	path := os.Args[1]

	var err error
	tasks, err = LoadTasksFromDir(path)
	if err != nil {
		panic(err)
	}

	ticker := time.Tick(1 * time.Second)

	go looper(tasks, ticker)
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
