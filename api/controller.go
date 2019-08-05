package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Controller handles the response to an HTTP request
type Controller struct {
	handler func(params map[string]string) (r interface{}, e error)
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

	response, err := c.handler(vars)
	if err != nil {
		if e, ok := err.(APIError); ok {
			w.WriteHeader(e.Status)
			w.Write(e.JSON())
		} else {
			log.WithField("err", err).Warn("Unexpected error occured")
			w.WriteHeader(500)
		}
		return
	}

	data, err := json.Marshal(struct {
		Data interface{} `json:"data"`
	}{
		Data: response,
	})

	if err != nil {
		apierr := NewAPIError(500, "Something failed in the system!")
		w.WriteHeader(apierr.Status)
		w.Write(apierr.JSON())
		return
	}

	w.WriteHeader(200)
	w.Write(data)
}

// NewController creates a Controller with passed in param func as its callback
func NewController(h func(params map[string]string) (r interface{}, e error)) *Controller {
	return &Controller{
		handler: h,
	}
}
