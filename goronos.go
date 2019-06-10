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

	// fmt.Println(string(data))

	var m Tasks

	err = yaml.UnmarshalStrict(data, &m)

	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}
