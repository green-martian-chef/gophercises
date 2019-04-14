package handler

import (
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
func YAMLHandler(yml string, fallback http.Handler) http.HandlerFunc {

	y := parseYAML(yml)
	p := buildMap(y)
	m := MapHandler(p, fallback)

	return m
}

func parseYAML(yml string) []pathToURL {
	var p []pathToURL

	err := yaml.Unmarshal([]byte(yml), &p)
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
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
