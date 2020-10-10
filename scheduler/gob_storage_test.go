package scheduler

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"
)

var tasksYaml = `name: command1
command: ["/path/to/command1"]
schedule:
  every: 2h30s
  months:
    - jun
    - 3
    - jul
    - dec
  weekdays:
    - sat
    - mon
    - tue
    - thu
  monthdays:
    - 13
    - 1
    - 18
  at:
    - 3:00
    - 5:13
    - 15:34
  except:
    weekdays:
      - thu
      - tue
    monthdays:
      - 13
    at:
      - 5
`

type gobTestSuite struct {
	suite.Suite

	tasks []*Task
}

type bufferCloser struct {
	bytes.Buffer
}

func (b *bufferCloser) Close() error {
	return nil
}

func (s *gobTestSuite) SetupSuite() {
	gob.Register(new(Task))
	gob.Register(new(Schedule))
	gob.Register(new(Hour))
	gob.Register(new(Between))

	storageMedium = new(bufferCloser)

	var task Task
	err := yaml.Unmarshal([]byte(tasksYaml), &task)
	if err != nil {
		panic(err)
	}
	s.tasks = make([]*Task, 1)
	s.tasks[0] = &task
}

func (s *gobTestSuite) TestStore() {
  store(s.tasks)

  var tasks []*Task
  load(&tasks)

  s.Assert().Equal(s.tasks, tasks)
}

func TestGob(t *testing.T) {
	suite.Run(t, new(gobTestSuite))
}
