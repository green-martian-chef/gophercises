package main

import (
	"net/http"
)

// MapHandler returns a handler function or a fallback
func MapHandler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if u, ok := pathsToURLs[path]; ok {
			http.Redirect(w, r, u, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}
