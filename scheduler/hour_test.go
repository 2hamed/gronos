package scheduler

import (
	"reflect"
	"testing"
)

func TestParseHour(t *testing.T) {
	h, err := parseHour("3")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(h, hour{hour: 3, minute: 0}) {
		t.Error("parseHour failed parsing single digit")
	}

	h, err = parseHour("3:15")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(h, hour{hour: 3, minute: 15}) {
		t.Error("parseHour failed parsing h:mm digit")
	}

	h, err = parseHour(":15")
	if err == nil {
		t.Error("err must not be nil here")
	}

	h, err = parseHour("30:25")
	if err == nil {
		t.Error("err must not be nil here")
	}

	h, err = parseHour("3:60")
	if err == nil {
		t.Error("err must not be nil here")
	}
}

func TestNewHour(t *testing.T) {
	_, err := NewHour(24, 0)
	if err == nil {
		t.Error("error must not be null")
	}

	_, err = NewHour(12, 61)
	if err == nil {
		t.Error("error must not be null")
	}
}
