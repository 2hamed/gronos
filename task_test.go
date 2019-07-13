package main

import (
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func TestTaskIsTime(t *testing.T) {
	var yamlStr = `
name: command1
command: /path/to/command1
schedule:
  every: "2:30"
  months:
    - jun
    - 3
    - dec
  weekdays:
    - sat
    - mon
    - tue
  monthdays:
    - 13
    - 1
  at:
    - 3:00
    - 5:13
    - 15:34
  except:
    every: "3:45"
    weekdays:
      - wed
      - tue
    monthdays:
      - 13
    at:
      - 5
`

	var task Task

	err := yaml.Unmarshal([]byte(yamlStr), &task)

	if err != nil {
		t.Error("the yaml is invalid", err)
	}

	anchor := time.Date(2019, time.June, 1, 3, 0, 0, 0, time.Local) // 1st jun 2019 3:0

	if !task.IsTime(&anchor) {
		t.Error("false negative")
	}

	anchor = time.Date(2019, time.June, 2, 3, 0, 0, 0, time.Local) // 2nd jun 2019 3:0

	if task.IsTime(&anchor) {
		t.Error("false positive")
	}

}
