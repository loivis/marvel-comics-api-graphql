package server

import (
	"net/http"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin == "" {
				origin = "*"
			}
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Headers", "*")
				// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Origin", origin)
				return
			}
			w.Header().Set("Access-Control-Allow-Origin", origin)
			h.ServeHTTP(w, r)
		},
	)
}
