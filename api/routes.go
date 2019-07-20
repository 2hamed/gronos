package api

import (
	"fmt"
	"net/http"

	"github.com/2hamed/goronos/scheduler"
	"github.com/gorilla/mux"
)

// RegisterRoutes regiters the api endpoints
func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/", homeHandler)
	router.Use(jsonMiddleWare)

}
func homeHandler(w http.ResponseWriter, r *http.Request) {

	output := []byte("")

	for _, v := range scheduler.GetTasks() {
		output = append(output, []byte(fmt.Sprintf("%v\n", v))...)
	}

	w.Write(output)
}
