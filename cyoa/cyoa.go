package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It Works! \n %s", r.URL.Path[1:])
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
