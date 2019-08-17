package scheduler

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	if !assert.Nil(t, err) {
		return
	}

	every := schedule.Every
	assert.Equal(t, 2*3600+30*60, every)

	weekdays := schedule.Weekdays

	assert.True(t, reflect.DeepEqual(weekdays, []time.Weekday{time.Saturday, time.Monday, time.Tuesday}))

	monthdays := schedule.Monthdays
	assert.True(t, reflect.DeepEqual(monthdays, []int{13}))
}

func TestScheduleAt(t *testing.T) {

	var schedule Schedule
	_ = yaml.Unmarshal([]byte(yamlStr), &schedule)

	anchor := time.Date(1, 1, 1, 5, 13, 0, 0, time.Local)

	a, _ := schedule.checkAt(&anchor)
	assert.True(t, a)

	anchor = time.Date(1, 1, 1, 5, 14, 0, 0, time.Local)

	a, _ = schedule.checkAt(&anchor)
	assert.False(t, a)
}

func TestScheduleMonths(t *testing.T) {
	var schedule Schedule
	_ = yaml.Unmarshal([]byte(yamlStr), &schedule)

	june := time.Date(2019, 6, 15, 5, 13, 0, 0, time.Local)

	m, _ := schedule.checkMonths(&june)
	assert.True(t, m)

	october := time.Date(2019, 10, 15, 5, 13, 0, 0, time.Local)

	m, _ = schedule.checkMonths(&october)
	assert.True(t, m)

	january := time.Date(2019, 1, 15, 5, 13, 0, 0, time.Local)

	m, _ = schedule.checkMonths(&january)
	assert.False(t, m)
}

func TestScheduleMothdays(t *testing.T) {
	var schedule Schedule
	_ = yaml.Unmarshal([]byte(yamlStr), &schedule)

	thirteenth := time.Date(2019, 6, 13, 5, 13, 0, 0, time.Local)

	twelfth := time.Date(2019, 6, 12, 5, 13, 0, 0, time.Local)

	md, _ := schedule.checkMonthdays(&thirteenth)
	assert.True(t, md)

	md, _ = schedule.checkMonthdays(&twelfth)
	assert.False(t, md)

}

func TestScheduleWeekdays(t *testing.T) {
	var schedule Schedule
	_ = yaml.Unmarshal([]byte(yamlStr), &schedule)

	saturday := time.Date(2019, 7, 13, 5, 13, 0, 0, time.Local) // 13th july

	wd, _ := schedule.checkWeekday(&saturday)
	assert.True(t, wd)

	friday := time.Date(2019, 7, 19, 5, 13, 0, 0, time.Local) // 19th july

	wd, _ = schedule.checkWeekday(&friday)
	assert.False(t, wd)

}
