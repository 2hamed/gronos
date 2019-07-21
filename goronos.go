package main

import (
	"fmt"
	"net/http"

	"github.com/2hamed/goronos/api"
	"github.com/2hamed/goronos/scheduler"
	"github.com/gorilla/mux"
)

func main() {

	scheduler.StartLooper()

	r := mux.NewRouter()

	api.RegisterRoutes(r)

	http.Handle("/", r)

	fmt.Println("Goronos engine is running and the API server is listening on port 8080...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
