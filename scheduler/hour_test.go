package scheduler

import (
	"reflect"
	"testing"
	"time"
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

	h, err = parseHour("5:a")
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

	h, err := NewHour(1, 34)
	if err != nil {
		t.Error("failed make new Hour")
	}
	if !reflect.DeepEqual(h, hour{hour: 1, minute: 34}) {
		t.Error("wrong hour created")
	}
}

func TestIsAfter(t *testing.T) {
	ti := time.Date(2019, 1, 1, 12, 12, 0, 0, time.Local)
	h, _ := NewHour(13, 9)

	if !h.IsAfter(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 13)

	if !h.IsAfter(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 10)

	if h.IsAfter(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(12, 10)

	if h.IsAfter(&ti) {
		t.Error("wrong behavior for IsAfter")
	}
}

func TestIsBefore(t *testing.T) {
	ti := time.Date(2019, 1, 1, 12, 12, 0, 0, time.Local)
	h, _ := NewHour(13, 9)

	if h.IsBefore(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 13)

	if h.IsBefore(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(13, 10)

	if !h.IsBefore(&ti) {
		t.Error("wrong behavior for IsAfter")
	}

	ti = time.Date(2019, 1, 1, 13, 12, 0, 0, time.Local)
	h, _ = NewHour(12, 10)

	if !h.IsBefore(&ti) {
		t.Error("wrong behavior for IsAfter")
	}
}
