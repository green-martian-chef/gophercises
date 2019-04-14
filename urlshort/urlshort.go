package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

func openFile(f *string) []byte {
	file, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatalf("The file could not be opened. \n%v", err)
	}
	return file
}

func main() {
	jsonFile := flag.String("json", "paths.json", "a JSON file containing an array of objects path:url")
	yamlFile := flag.String("yaml", "paths.yaml", "a YAML file containing a list of path:url")
	isJSON := flag.Bool("j", false, "use an JSON file instead of a YAML file. Use -json to provide a file.")
	isYAML := flag.Bool("y", false, "use an YAML file instead the default map. Use -yaml to provide a file.")
	flag.Parse()

	m := map[string]string{
		"/godoc-http":     "https://golang.org/pkg/net/http/",
		"/godoc-yaml":     "https://godoc.org/gopkg.in/yaml.v2",
		"/urlshort":       "https://github.com/gophercises/urlshort",
		"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
	}

	r := defaultHandler()
	var h http.HandlerFunc

	// Default to map of paths and URLs
	if *isJSON == true {
		file := openFile(jsonFile)
		h = handler.JSONHandler(file, r)
	} else if *isYAML == true {
		file := openFile(yamlFile)
		h = handler.YAMLHandler(file, r)
	} else {
		h = handler.MapHandler(m, r)
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
