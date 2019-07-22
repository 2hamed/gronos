package scheduler

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseHour(t *testing.T) {
	h, err := parseHour("3")
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(h, hour{hour: 3, minute: 0}))

	h, err = parseHour("3:15")
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(h, hour{hour: 3, minute: 15}))

	h, err = parseHour(":15")
	assert.NotNil(t, err)

	h, err = parseHour("5:a")
	assert.NotNil(t, err)

	h, err = parseHour("30:25")
	assert.NotNil(t, err)

	h, err = parseHour("3:60")
	assert.NotNil(t, err)

}

func TestNewHour(t *testing.T) {
	_, err := NewHour(24, 0)
	assert.NotNil(t, err)

	_, err = NewHour(12, 61)
	assert.NotNil(t, err)

	h, err := NewHour(1, 34)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(h, hour{hour: 1, minute: 34}))
}

func TestIsAfter(t *testing.T) {
	ti := time.Date(2019, 1, 1, 12, 12, 0, 0, time.Local)
	h, _ := NewHour(13, 9)
	assert.True(t, h.IsAfter(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 13)
	assert.True(t, h.IsAfter(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 10)
	assert.False(t, h.IsAfter(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(12, 10)
	assert.False(t, h.IsAfter(&ti))

}

func TestIsBefore(t *testing.T) {
	ti := time.Date(2019, 1, 1, 12, 12, 0, 0, time.Local)
	h, _ := NewHour(13, 9)
	assert.False(t, h.IsBefore(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 13)
	assert.False(t, h.IsBefore(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 10)
	assert.True(t, h.IsBefore(&ti))

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(12, 10)
	assert.True(t, h.IsBefore(&ti))

}
