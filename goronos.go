package main

import (
	"fmt"
	"os"
	"time"
)

var tasks Tasks

func main() {
	path := os.Args[1]

	tasks, err := LoadTasksFromDir(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)

	ticker := time.Tick(5 * time.Second)
	forever := make(chan struct{})

	go func() {
		for {
			t := <-ticker
			for taskName, task := range tasks {
				if task.Schedule.IsTime(taskName, t) {
					task.Execute()
					taskLastRunTime[taskName] = time.Now().Unix()
				}
			}
		}
	}()

	<-forever
}
