package scheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tasks = []*Task{
	&Task{
		Name: "task1",
	},
	&Task{
		Name: "task2",
	},
	&Task{
		Name: "task3",
	},
}

func TestGetTasks(t *testing.T) {

	tm.initialize(tasks)

	ts := GetTasks()
	assert.ElementsMatch(t, ts, tasks)
}

func TestGetTask(t *testing.T) {

	tm.initialize(tasks)

	ts, err := GetTask("task2")
	assert.Nil(t, err)

	assert.Equal(t, ts, tasks[1])
}

func TestDisableTask(t *testing.T) {
	tm.initialize(tasks)

	err := DisableTask("task")

	assert.NotNil(t, err)

	err = DisableTask("task1")

	assert.Nil(t, err)

	ts := GetTasks()

	assert.ElementsMatch(t, ts, []*Task{
		&Task{
			Name: "task2",
		},
		&Task{
			Name: "task3",
		},
	})

}

func TestEnableTask(t *testing.T) {
	tm.initialize(tasks)

	err := DisableTask("task1")
	assert.Nil(t, err)

	err = EnableTask("task")

	assert.NotNil(t, err)

	err = EnableTask("task1")

	assert.Nil(t, err)

	assert.ElementsMatch(t, GetTasks(), tasks)
}

func TestGetDisabledTask(t *testing.T) {
	tm.initialize(tasks)

	err := DisableTask("task1")
	assert.Nil(t, err)

	err = DisableTask("task2")
	assert.Nil(t, err)

	assert.ElementsMatch(t, GetDisabledTasks(), []*Task{
		&Task{
			Name: "task1",
		},
		&Task{
			Name: "task2",
		},
	})

	err = EnableTask("task2")
	assert.Nil(t, err)

	assert.ElementsMatch(t, GetDisabledTasks(), []*Task{
		&Task{
			Name: "task1",
		},
	})
}
