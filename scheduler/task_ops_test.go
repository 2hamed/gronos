package scheduler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var taskss = []*Task{
	{
		Name: "task1",
	},
	{
		Name: "task2",
	},
	{
		Name: "task3",
	},
}

func TestGetTasks(t *testing.T) {

	tm.initialize(taskss)

	ts := GetTasks()
	assert.ElementsMatch(t, ts, taskss)
}

func TestGetTask(t *testing.T) {

	tm.initialize(taskss)

	ts, err := GetTask("task2")
	assert.Nil(t, err)

	assert.Equal(t, ts, taskss[1])
}

func TestDisableTask(t *testing.T) {
	tm.initialize(taskss)

	err := DisableTask("task")

	assert.NotNil(t, err)

	err = DisableTask("task1")

	assert.Nil(t, err)

	ts := GetTasks()

	assert.ElementsMatch(t, ts, []*Task{
		{
			Name: "task2",
		},
		{
			Name: "task3",
		},
	})

}

func TestEnableTask(t *testing.T) {
	tm.initialize(taskss)

	err := DisableTask("task1")
	assert.Nil(t, err)

	err = EnableTask("task")

	assert.NotNil(t, err)

	err = EnableTask("task1")

	assert.Nil(t, err)

	assert.ElementsMatch(t, GetTasks(), taskss)
}

func TestGetDisabledTask(t *testing.T) {
	tm.initialize(taskss)

	err := DisableTask("task1")
	assert.Nil(t, err)

	err = DisableTask("task2")
	assert.Nil(t, err)

	assert.ElementsMatch(t, GetDisabledTasks(), []*Task{
		{
			Name: "task1",
		},
		{
			Name: "task2",
		},
	})

	err = EnableTask("task2")
	assert.Nil(t, err)

	assert.ElementsMatch(t, GetDisabledTasks(), []*Task{
		{
			Name: "task1",
		},
	})
}
