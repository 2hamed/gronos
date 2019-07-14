package scheduler

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
 - sat
 - mon
 - tue
monthdays:
 - 13
at:
 - 3:00
 - 5:13
 - 15:34
months:
 - jun
 - 3
 - october
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

	if !reflect.DeepEqual(weekdays, []time.Weekday{time.Saturday, time.Monday, time.Tuesday}) {
		t.Error("wrong value for `weekdays`")
	}

	monthdays := schedule.Monthdays()

	if !reflect.DeepEqual(monthdays, []int{13}) {
		t.Error("wrong value for `monthdays`")
	}
}

func TestScheduleAt(t *testing.T) {

	var schedule Schedule
	yaml.Unmarshal([]byte(yamlStr), &schedule)

	anchor := time.Date(1, 1, 1, 5, 13, 0, 0, time.Local)

	if !schedule.checkAt(&anchor) {
		t.Error("checkAt failed to recognize")
	}

	anchor = time.Date(1, 1, 1, 5, 14, 0, 0, time.Local)

	if schedule.checkAt(&anchor) {
		t.Error("checkAt failed to recognize")
	}
}

func TestScheduleMonths(t *testing.T) {
	var schedule Schedule
	yaml.Unmarshal([]byte(yamlStr), &schedule)

	june := time.Date(2019, 6, 15, 5, 13, 0, 0, time.Local)

	if !schedule.checkMonths(&june) {
		t.Error("failed to recognize June")
	}

	october := time.Date(2019, 10, 15, 5, 13, 0, 0, time.Local)

	if !schedule.checkMonths(&october) {
		t.Error("failed to recognize October")
	}

	january := time.Date(2019, 1, 15, 5, 13, 0, 0, time.Local)

	if schedule.checkMonths(&january) {
		t.Error("false positive in months")
	}
}

func TestScheduleMothdays(t *testing.T) {
	var schedule Schedule
	yaml.Unmarshal([]byte(yamlStr), &schedule)

	thirteenth := time.Date(2019, 6, 13, 5, 13, 0, 0, time.Local)

	twelfth := time.Date(2019, 6, 12, 5, 13, 0, 0, time.Local)

	if !schedule.checkMonthdays(&thirteenth) {
		t.Error("failed recognizing 13th")
	}

	if schedule.checkMonthdays(&twelfth) {
		t.Error("false positive in recognizing 12th")
	}
}

func TestScheduleWeekdays(t *testing.T) {
	var schedule Schedule
	yaml.Unmarshal([]byte(yamlStr), &schedule)

	saturday := time.Date(2019, 7, 13, 5, 13, 0, 0, time.Local) // 13th july

	if !schedule.checkWeekday(&saturday) {
		t.Error("failed to recognize saturday")
	}

	friday := time.Date(2019, 7, 19, 5, 13, 0, 0, time.Local) // 19th july

	if schedule.checkWeekday(&friday) {
		t.Error("false positive in recognizing weekday")
	}

}
