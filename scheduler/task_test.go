package scheduler

import (
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

var taskYaml = `
name: command1
command: ["/path/to/command1"]
schedule:
  every: "2:30"
  months:
    - jun
    - 3
    - jul
    - dec
  weekdays:
    - sat
    - mon
    - tue
    - thu
  monthdays:
    - 13
    - 1
    - 18
  at:
    - 3:00
    - 5:13
    - 15:34
  except:
    weekdays:
      - thu
      - tue
    monthdays:
      - 13
    at:
      - 5
`

func TestTaskIsTime(t *testing.T) {
	var task Task

	err := yaml.Unmarshal([]byte(taskYaml), &task)

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

	// testing the Except part
	anchor = time.Date(2019, time.July, 18, 3, 0, 0, 0, time.Local) // 18th jul 2019 3:0

	if task.IsTime(&anchor) {
		t.Error("false positive, this should be excepted")
	}

}
