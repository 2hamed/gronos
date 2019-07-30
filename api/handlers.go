package api

import (
	"encoding/json"
	"log"

	"github.com/2hamed/goronos/scheduler"
)

func listTasks(params map[string]string) ([]byte, *APIError) {

	output, _ := json.Marshal(scheduler.GetTasks())

	return output, nil
}

func getTask(params map[string]string) ([]byte, *APIError) {

	log.Println(params)

	return nil, nil
}
