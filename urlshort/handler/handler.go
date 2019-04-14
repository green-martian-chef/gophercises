package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
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

// YAMLHandler returns a handler or a fallback
func YAMLHandler(yml []byte, fallback http.Handler) http.HandlerFunc {

	y := parseYAML(yml)
	p := buildMap(y)
	m := MapHandler(p, fallback)

	return m
}

// JSONHandler returns a handler or a fallback
func JSONHandler(j []byte, fallback http.Handler) http.HandlerFunc {
	jp := parseJSON(j)
	p := buildMap(jp)
	m := MapHandler(p, fallback)

	return m
}

func parseYAML(yml []byte) []pathToURL {
	var p []pathToURL

	err := yaml.Unmarshal(yml, &p)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return p
}

func parseJSON(j []byte) []pathToURL {
	var p []pathToURL

	err := json.Unmarshal(j, &p)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return p
}
func buildMap(p []pathToURL) map[string]string {
	pathsToURLs := make(map[string]string)

	for _, v := range p {
		pathsToURLs[v.Path] = v.URL
	}

	return pathsToURLs
}

// pathToURL represents a path and its corresponding URL
type pathToURL struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url"`
}
