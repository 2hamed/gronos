package scheduler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// NewHour creates an Hour
func NewHour(h, m int) (Hour, error) {
	hour:=Hour{}
	if h > 23 || h < 0 {
		return hour, errors.New("h can not be less than 0 or more than 23")
	}
	if m > 59 || m < 0 {
		return hour, errors.New("m can not be less than 0 or more than 59")
	}
	return Hour{
		Hour:   h,
		Minute: m,
	}, nil
}

func newHourNoErr(h, m int) Hour {
	hour, _ := NewHour(h, m)
	return hour
}

type Hour struct {
	Hour   int
	Minute int
}

func (h Hour) String() string {
	return fmt.Sprintf("%d:%02d", h.Hour, h.Minute)
}

func (h Hour) MarshalJSON() ([]byte, error) {
	data := make(map[string]int)

	data["hour"] = h.Hour
	data["minute"] = h.Minute

	return json.Marshal(data)
}

func (h Hour) IsAfter(t *time.Time) bool {
	if h.Hour == t.Hour() && h.Minute >= t.Minute() {
		return true
	}

	if h.Hour > t.Hour() {
		return true
	}

	return false
}

func (h Hour) IsBefore(t *time.Time) bool {
	if h.Hour == t.Hour() && h.Minute <= t.Minute() {
		return true
	}

	if h.Hour < t.Hour() {
		return true
	}
	return false
}


// parseHour parses a string in the format of HH:mm into an `hour` struct
func parseHour(str string) (Hour, error) {
	hm := strings.Split(str, ":")
	h, err := strconv.Atoi(hm[0])

	if err != nil {
		return Hour{}, errors.Wrap(err, "failed parsing the hour")
	}
	var m int
	if len(hm) > 1 {
		m, err = strconv.Atoi(hm[1])
		if err != nil {
			return Hour{}, errors.Wrap(err, "failed parsing the minute")
		}
	}

	return NewHour(h, m)
}
