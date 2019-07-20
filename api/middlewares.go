package api

import (
	"net/http"
)

func jsonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Author", "Hamed Momeni")

		next.ServeHTTP(w, r)

	})
}
