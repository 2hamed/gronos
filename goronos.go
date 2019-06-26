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

	// fmt.Println(tasks)

	ticker := time.Tick(5 * time.Second)
	forever := make(chan struct{})

	go func() {
		for {
			_ = <-ticker
			fmt.Println("checking tasks")
			for _, task := range tasks {
				if task.IsTime() {
					task.Execute()
				}
			}
		}
	}()

	fmt.Println("Goronos is now active...")
	<-forever
}
