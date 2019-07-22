package scheduler

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Between represents a period in a day
type Between struct {
	From Hour `json:"from"`
	To   Hour `json:"to"`
}

func (b Between) String() string {
	return fmt.Sprintf("%s-%s", b.From.String(), b.To.String())
}

// IsInside checks whether a given time.Time is inside the period or not
func (b Between) IsInside(t *time.Time) bool {
	return b.From.IsBefore(t) && b.To.IsAfter(t)
}

func parseBetween(str string) (Between, error) {
	parts := strings.Split(str, "-")
	if len(parts) < 2 {
		return Between{}, errors.New("a between statement must have both a start and an end")
	}

	from, err := parseHour(parts[0])

	if err != nil {
		return Between{}, errors.Wrap(err, "failed parsing the between statement")
	}

	to, err := parseHour(parts[1])

	if err != nil {
		return Between{}, errors.Wrap(err, "failed parsing the between statement")
	}

	return Between{
		From: from,
		To:   to,
	}, nil
}
