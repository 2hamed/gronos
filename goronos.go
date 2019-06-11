package main

import (
	"fmt"
	"os"
)

var tasks Tasks

func main() {
	path := os.Args[1]

	tasks, err := LoadTasksFromDir(path)
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks["task1"].Schedules[0].At())
}
