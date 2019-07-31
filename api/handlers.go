package api

import (
	"github.com/2hamed/goronos/scheduler"
)

func listTasks(params map[string]string) (*Response, *APIError) {

	return &Response{Data: scheduler.GetTasks()}, nil
}

func getTask(params map[string]string) (*Response, *APIError) {
	task, err := scheduler.GetTask(params["name"])
	if err != nil {
		return nil, NewAPIError(HTTP_NOT_FOUND, "Task with that name not found!")
	}

	return &Response{Data: task}, nil
}
