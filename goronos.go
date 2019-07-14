package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/2hamed/goronos/scheduler"
	"github.com/gorilla/mux"
)

var tasks scheduler.Tasks

func main() {
	path := os.Args[1]

	var err error

	tasks, err = scheduler.LoadTasksFromDir(path)
	if err != nil {
		panic(err)
	}

	ticker := time.Tick(5 * time.Second)

	go looper(tasks, ticker)

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func looper(tasks scheduler.Tasks, ticker <-chan time.Time) {
	for {
		t := <-ticker
		for _, task := range tasks {
			if task.IsTime(&t) {
				task.Execute()
			}
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	output := []byte("")

	for _, v := range tasks {
		output = append(output, []byte(fmt.Sprintf("%v\n", v))...)
	}

	w.Write(output)
}
