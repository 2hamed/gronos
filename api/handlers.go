package api

import (
	"encoding/json"

	"github.com/2hamed/goronos/scheduler"
)

func listTasks(params map[string]string) ([]byte, *APIError) {

	output, err := json.Marshal(scheduler.GetTasks())

	if err != nil {
		return nil, NewAPIError(HTTP_SERVER_ERROR, err.Error(), err)
	}

	return output, nil
}

func getTask(params map[string]string) ([]byte, *APIError) {

	return nil, NewAPIError(HTTP_NOT_FOUND, "Task with that name not found!")
}
