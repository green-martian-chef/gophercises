package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type story map[string]chapter

type chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

var tmpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>{{.Title}}</title>
  </head>
	<body>
		<main>
			<section>
				<header>
					<h1>{{.Title}}</h1>
				</header>
				<article>
					{{range .Story}}
						<p>{{.}}</p>
					{{end}}
					<hr>
					<ul>
					{{range .Options}}
					<li>
						<a href="/{{.Arc}}">{{.Text}}</a>
					</li>
					{{end}}
				</article>
			</section>
		</main>
  </body>
</html>
`

func handler(j []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]
		if len(path) < 1 {
			path = "intro"
		}

		chapter := &path

		p, err := parseJSON(j)
		if err != nil {
			log.Fatalf("JSON couldn't be parsed %v", err)
		}

		t, err := template.New("tmpl").Parse(tmpl)
		if err != nil {
			log.Fatalf("Template couldn't be parsed %v", err)
		}

		err = t.Execute(w, p[*chapter])
		if err != nil {
			log.Fatalf("Template execution failed %v", err)
		}

	}
}

func parseJSON(j []byte) (story, error) {
	var ret story

	if err := json.Unmarshal(j, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func main() {
	jsonFile, err := ioutil.ReadFile("gopher.json")

	if err != nil {
		log.Fatalf("Failed to open JSON file \n %v", err)
	}

	h := handler(jsonFile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", h)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
