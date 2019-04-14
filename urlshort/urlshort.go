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
	yml := `
- path: /godoc-http
  url: https://golang.org/pkg/net/http/
- path: /godoc-yaml
  url: https://godoc.org/gopkg.in/yaml.v2
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	r := defaultHandler()
	y := YAMLHandler(yml, r)

	s := &http.Server{
		Addr:    ":8080",
		Handler: y,
	}

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
