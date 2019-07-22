package scheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	assert.Nil(t, err, "the yaml is invalid")

	anchor := time.Date(2019, time.June, 1, 3, 0, 0, 0, time.Local) // 1st jun 2019 3:0
	assert.True(t, task.IsTime(&anchor))

	anchor = time.Date(2019, time.June, 2, 3, 0, 0, 0, time.Local) // 2nd jun 2019 3:0
	assert.False(t, task.IsTime(&anchor))

	// testing the Except part
	anchor = time.Date(2019, time.July, 13, 3, 0, 0, 0, time.Local) // 13th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))

	anchor = time.Date(2019, time.July, 18, 3, 0, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))

}

func TestExceptAt(t *testing.T) {
	var task Task

	yamlStr = `
name: command1
command: ["/path/to/command1"]
schedule:
  every: "1"
  except:
    at:
      - 5:15`
	err := yaml.Unmarshal([]byte(yamlStr), &task)
	assert.Nil(t, err)

	anchor := time.Date(2019, time.July, 18, 5, 15, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))
}

func TestExceptMonths(t *testing.T) {
	var task Task

	yamlStr = `
name: command1
command: ["/path/to/command1"]
schedule:
  at:
    - 5:15
  except:
    months:
      - jul`
	err := yaml.Unmarshal([]byte(yamlStr), &task)
	assert.Nil(t, err, "the yaml is invalid")

	anchor := time.Date(2019, time.July, 18, 5, 15, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))
}

func TestBetweens(t *testing.T) {
	var task Task

	yamlStr = `
name: command1
command: ["/path/to/command1"]
schedule:
  between:
    - 5-6:30
  except:
    at:
      - 5:15
    between:
      - 6:15-7`
	err := yaml.Unmarshal([]byte(yamlStr), &task)
	assert.Nil(t, err, "the yaml is invalid")

	anchor := time.Date(2019, time.July, 18, 5, 30, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.True(t, task.IsTime(&anchor))

	anchor = time.Date(2019, time.July, 18, 5, 15, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))

	anchor = time.Date(2019, time.July, 18, 6, 25, 0, 0, time.Local) // 18th jul 2019 3:0
	assert.False(t, task.IsTime(&anchor))

}

func TestRepeatingTask(t *testing.T) {
	var task Task

	yamlStr = `
name: command1
command: ["echo", "hello", ">", "/dev/null"]
schedule:
  every: 0:30`
	err := yaml.Unmarshal([]byte(yamlStr), &task)

	assert.Nil(t, err)

	now := time.Now()

	task.Execute()

	now = now.Add(5 * time.Minute)
	assert.False(t, task.IsTime(&now))

	now = now.Add(30 * time.Minute)
	assert.True(t, task.IsTime(&now))
}
