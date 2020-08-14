package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/2hamed/goronos/api"
	"github.com/2hamed/goronos/scheduler"
	"github.com/2hamed/goronos/storage"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {

	logFile, err := os.OpenFile("./logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0755))

	if err != nil {
		panic(err)
	}

	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(logFile)

	log.SetLevel(log.DebugLevel)
}
func main() {

	configPath := os.Args[1]

	storage, err := storage.NewStorage("./store")
	if err != nil {
		panic(err)
	}

	scheduler.SetStorage(storage)
	scheduler.StartLooper(configPath)

	r := mux.NewRouter()

	api.RegisterRoutes(r)

	http.Handle("/", r)

	fmt.Println("Gronos engine is running and the API server is listening on port 8080...")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
