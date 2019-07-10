package main

import (
	"testing"

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
    - wed
    - mon
    - tue
  monthdays:
    - 13
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

  
}
