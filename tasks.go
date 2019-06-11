package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"syscall"

	"gopkg.in/yaml.v2"
)

// Tasks is a type alias of map[string]Task
type Tasks map[string]Task

// Task is command which run by Schdule
type Task struct {
	Command  string   `yaml:"command"`
	Schedule Schedule `yaml:"schedule"`
}

// LoadTasksFromFile reads an etire YAML file and outputs the corresponding Tasks struct
func LoadTasksFromFile(filePath string) (Tasks, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

// LoadTasksFromDir scans a directory and loads every YAML file into corresponding Tasks struct
func LoadTasksFromDir(dirPath string) (Tasks, error) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var tasks = make(Tasks)

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yaml") {
			t, err := LoadTasksFromFile(dirPath + f.Name())
			if err != nil {
				panic(err)
			}

			for k, task := range t {
				if _, ok := tasks[k]; ok {
					return nil, errors.New("duplicate task name: " + k)
				}
				tasks[k] = task
			}
		}
	}

	return tasks, nil
}

func (task Task) Execute() {
	go func() {
		err := syscall.Exec(task.Command, nil, nil)
		if err != nil {
			log.Println(fmt.Errorf("Running command %s failed: %v", task.Command, err))
		}
	}()
}
