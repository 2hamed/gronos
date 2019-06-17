package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var yamlStr = `
every: 2:30
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

func TestScheduleUnmarshal(t *testing.T) {

	var schedule Schedule
	err := yaml.Unmarshal([]byte(yamlStr), &schedule)

	if err != nil {
		t.Error(errors.Wrap(err, "failed to unmarshal yaml"))
		return
	}

	every, err := schedule.Every()

	if every != 2*3600+30*60 || err != nil {
		t.Error(errors.Wrap(err, "wrong value for `every`"))
	}

	weekdays := schedule.Weekdays()

	if !reflect.DeepEqual(weekdays, []time.Weekday{time.Wednesday, time.Monday, time.Tuesday}) {
		t.Error("wrong value for `weekdays`")
	}

	monthdays := schedule.monthdays

	if !reflect.DeepEqual(monthdays, []int{13}) {
		t.Error("wrong value for `monthdays`")
	}
}
