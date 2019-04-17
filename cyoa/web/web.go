package web

import (
	"log"
	"net/http"
	"text/template"

	"github.com/tscangussu/gophercises/cyoa/parser"
)

// Handler returns a HandlerFunc
func Handler(s parser.Story, t string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if len(path) < 1 {
			path = "intro"
		}

		chapter := &path

		t, err := template.New("tmpl").Parse(t)
		if err != nil {
			log.Fatalf("Template couldn't be parsed %v", err)
		}

		err = t.Execute(w, s[*chapter])
		if err != nil {
			log.Fatalf("Template execution failed %v", err)
		}

	}
}

// Server creates a new mux and returns a Server
func Server(h http.HandlerFunc) *http.Server {
	m := http.NewServeMux()
	m.HandleFunc("/", h)

	s := &http.Server{
		Addr:    ":8080",
		Handler: m,
	}

	return s
}
