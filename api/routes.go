package api

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes regiters the api endpoints
func RegisterRoutes(router *mux.Router) {

	router.Handle("/tasks", NewController(listTasks)).Methods("GET")
	router.Handle("/tasks/disabled", NewController(getDisabledTasks)).Methods("GET")
	router.Handle("/tasks/{name}", NewController(getTask)).Methods("GET")
	router.Handle("/tasks/{name}/disable", NewController(disableTask)).Methods("GET")
	router.Handle("/tasks/{name}/enable", NewController(enableTask)).Methods("GET")

	router.Use(jsonMiddleWare)

}
