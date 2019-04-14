package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tscangussu/gophercises/urlshort/handler"
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
	isJSON := flag.Bool("j", false, "use JSON instead YAML")
	flag.Parse()

	jsonPaths := `
	[
		{
			"path": "/godoc-http",
			"url": "https://golang.org/pkg/net/http/"
		},
		{
			"path": "/godoc-yaml",
			"url": "https://godoc.org/gopkg.in/yaml.v2"
		},
		{
			"path": "/urlshort",
			"url": "https://github.com/gophercises/urlshort"
		},
		{
			"path": "/urlshort-final",
			"url": "https://github.com/gophercises/urlshort/tree/solution"
		}
	]	
`

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
	var h http.HandlerFunc

	if *isJSON == true {
		h = handler.JSONHandler(jsonPaths, r)
	} else {
		h = handler.YAMLHandler(yml, r)

	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
