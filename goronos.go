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

	m := make(map[interface{}]interface{})

	yaml.Unmarshal(data, &m)

	fmt.Println(m)
}
