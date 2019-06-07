package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	path := os.Args[1]

	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	m := make(map[string]Task)

	err = yaml.Unmarshal(data, &m)

	if err != nil {
		panic(err)
	}

	fmt.Println(m["task1"].Schedules)
}
