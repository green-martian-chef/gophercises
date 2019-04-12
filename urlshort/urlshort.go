package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler() (r *http.ServeMux) {
	r = http.NewServeMux()
	r.HandleFunc("/", home)
	return
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It Works!")
}

func main() {
	pathsToURLs := map[string]string{
		"/godoc-http": "https://golang.org/pkg/net/http/",
		"/godoc-yaml": "https://godoc.org/gopkg.in/yaml.v2",
	}

	r := defaultHandler()
	m := MapHandler(pathsToURLs, r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}
	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
