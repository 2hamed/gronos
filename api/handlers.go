package api

import (
	sch "github.com/2hamed/goronos/scheduler"
)

func listTasks(params map[string]string) (interface{}, error) {

	return sch.GetTasks(), nil
}

func getTask(params map[string]string) (interface{}, error) {
	task, err := sch.GetTask(params["name"])
	if err != nil {
		return nil, NewAPIError(HTTP_NOT_FOUND, "Task with that name not found!")
	}

	return task, nil
}

func getDisabledTasks(params map[string]string) (interface{}, error) {
	return sch.GetDisabledTasks(), nil
}

func disableTask(params map[string]string) (interface{}, error) {
	name := params["name"]

	err := sch.DisableTask(name)

	if err != nil {
		return nil, NewAPIError(HTTP_NOT_FOUND, err.Error())
	}

	return getDisabledTasks(params)
}
func enableTask(params map[string]string) (interface{}, error) {
	name := params["name"]

	err := sch.EnableTask(name)

	if err != nil {
		return nil, NewAPIError(HTTP_NOT_FOUND, err.Error())
	}

	return getDisabledTasks(params)
}
