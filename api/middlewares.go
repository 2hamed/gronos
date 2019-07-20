package api

import (
	"context"
	"net/http"
)

type Key string

var key Key = "KEY"

func jsonMiddleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Author", "Hamed Momeni")

		r = r.WithContext(context.WithValue(r.Context(), key, "VAL"))

		next.ServeHTTP(w, r)

	})
}
