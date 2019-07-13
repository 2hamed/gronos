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

	ticker := time.Tick(5 * time.Second)
	forever := make(chan struct{})

	go looper(tasks, ticker)

	fmt.Println("Goronos is now active...")
	<-forever
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
