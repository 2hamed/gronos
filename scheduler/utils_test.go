package scheduler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntSliceContains(t *testing.T) {
	haystack := []int{1, 3, 4, 6, 712, 45, -78}

	assert.True(t, IntSliceContains(haystack, 4))

	assert.False(t, IntSliceContains(haystack, 5))

	assert.False(t, IntSliceContains(nil, 5))
}

func TestWeekdaySliceContains(t *testing.T) {
	haystack := []time.Weekday{time.Thursday, time.Wednesday}

	assert.False(t, WeekDaySliceContains(haystack, time.Saturday))

	assert.True(t, WeekDaySliceContains(haystack, time.Thursday))

	assert.False(t, WeekDaySliceContains(nil, time.Saturday))
}
func TestMonthSliceContains(t *testing.T) {
	haystack := []time.Month{time.January, time.March, time.October}

	if MonthSliceContains(haystack, time.June) {
		t.Error("haystack does not contain time.June")
	}

	if !MonthSliceContains(haystack, time.March) {
		t.Error("haystack does contain time.March")
	}

	if MonthSliceContains(nil, time.January) {
		t.Error("nil slice should always return false")
	}
}

func TestTimeSliceContainsHoursMintues(t *testing.T) {
	haystack := []time.Time{
		time.Date(0, 0, 0, 13, 13, 0, 0, time.Local),
		time.Date(0, 0, 0, 3, 0, 0, 0, time.Local),
		time.Date(2010, 1, 1, 19, 1, 0, 0, time.Local),
		time.Date(0, 0, 0, 5, 59, 0, 0, time.Local),
	}

	assert.True(t, TimeSliceContainsHoursMintues(haystack, time.Date(4, 4, 4, 13, 13, 10, 0, time.Local)))

	assert.False(t, TimeSliceContainsHoursMintues(haystack, time.Date(2010, 1, 1, 19, 0, 0, 0, time.Local)))

	assert.False(t, TimeSliceContainsHoursMintues(nil, time.Date(4, 5, 6, 73, 4, 6, 4, time.Local)))
}

func TestParseMonth(t *testing.T) {
	m, err := ParseMonth(1)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, time.January, m)

	m, err = ParseMonth("jun")
	assert.Nil(t, err)
	assert.Equal(t, time.June, m)

	m, err = ParseMonth("january")
	assert.Nil(t, err)
	assert.Equal(t, time.January, m)

	_, err = ParseMonth("")
	assert.NotNil(t, err)
}
