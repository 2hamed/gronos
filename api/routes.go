package api

import (
	"encoding/json"
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
	fmt.Println(r.Context().Value(key))

	output, _ := json.Marshal(scheduler.GetTasks())

	w.Write(output)
}
