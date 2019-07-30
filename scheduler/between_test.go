package scheduler

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseBetween(t *testing.T) {
	b, err := parseBetween("3-5")
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(b, Between{From: newHourNoErr(3, 0), To: newHourNoErr(5, 0)}))

	b, err = parseBetween("17-18:40")
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(b, Between{From: newHourNoErr(17, 0), To: newHourNoErr(18, 40)}))

	b, err = parseBetween("35")
	assert.NotNil(t, err)

	b, err = parseBetween("3:70-6")
	assert.NotNil(t, err)

}

func TestIsInside(t *testing.T) {
	b, _ := parseBetween("3:30-5")
	tim := time.Date(2019, 1, 1, 4, 0, 0, 0, time.Local)
	assert.True(t, b.IsInside(&tim))

	b, _ = parseBetween("3:30-5")
	tim = time.Date(2019, 1, 1, 5, 1, 0, 0, time.Local)
	assert.False(t, b.IsInside(&tim))

}
