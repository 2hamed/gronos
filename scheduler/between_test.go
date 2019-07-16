package scheduler

import (
	"reflect"
	"testing"
	"time"
)

func TestParseBetween(t *testing.T) {
	b, err := parseBetween("3-5")
	if err != nil {
		t.Error("failed parsing valid between")
	}

	if !reflect.DeepEqual(b, Between{from: newHourNoErr(3, 0), to: newHourNoErr(5, 0)}) {
		t.Error("wrong parsed value for between")
	}

	b, err = parseBetween("17-18:40")
	if err != nil {
		t.Error("failed parsing valid between")
	}

	if !reflect.DeepEqual(b, Between{from: newHourNoErr(17, 0), to: newHourNoErr(18, 40)}) {
		t.Error("wrong parsed value for between")
	}

	b, err = parseBetween("35")
	if err == nil {
		t.Error("should not be able to parse this!")
	}

	b, err = parseBetween("3:70-6")
	if err == nil {
		t.Error("should not be able to parse this!")
	}

}

func TestIsInside(t *testing.T) {
	b, _ := parseBetween("3:30-5")
	tim := time.Date(2019, 1, 1, 4, 0, 0, 0, time.Local)

	if !b.IsInside(&tim) {
		t.Error("wrong inside detection")
	}

	b, _ = parseBetween("3:30-5")
	tim = time.Date(2019, 1, 1, 5, 1, 0, 0, time.Local)

	if b.IsInside(&tim) {
		t.Error("wrong inside detection")
	}
}
