package api

import (
	"encoding/json"
	"log"

	"github.com/2hamed/goronos/scheduler"
)

func listTasks(params map[string]string) []byte {

	output, _ := json.Marshal(scheduler.GetTasks())

	return output
}

func getTask(params map[string]string) []byte {

	log.Println(params)

	return nil
}
