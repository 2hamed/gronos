package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Controller handles the response to an HTTP request
type Controller struct {
	handler func(params map[string]string) (r []byte, e *APIError)
}

func (c Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	vals := r.Form

	for k, v := range vals {
		if len(v) == 0 {
			continue
		}
		vars[k] = v[0]
	}

	data, err := c.handler(vars)

	if err != nil {
		w.WriteHeader(err.Status)
		w.Write(err.JSON())
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

// NewController creates a Controller with passed in param func as its callback
func NewController(h func(params map[string]string) (r []byte, e *APIError)) *Controller {
	return &Controller{
		handler: h,
	}
}
