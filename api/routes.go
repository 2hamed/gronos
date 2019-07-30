package api

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes regiters the api endpoints
func RegisterRoutes(router *mux.Router) {

	router.Handle("/", NewController(listTasks)).Methods("GET")
	router.Handle("/{name}", NewController(getTask)).Methods("GET")

	router.Use(jsonMiddleWare)

}
