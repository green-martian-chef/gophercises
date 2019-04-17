package main

import (
	"io/ioutil"
	"log"

	"github.com/tscangussu/gophercises/cyoa/parser"
	"github.com/tscangussu/gophercises/cyoa/web"
)

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
					<ul>
					{{range .Options}}
					<li>
						<a href="/{{.Arc}}">{{.Text}}</a>
					</li>
					{{end}}
				</article>
			</section>
		</main>
		<style>
		*, ::before, ::after { box-sizing: border-box; } /* Switch to border-box for box-sizing. */
		<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      section {
        width: 80%;
        max-width: 700px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #f7fdff;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #5086a9;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
		@media only screen and (min-width : 576px) {}
		@media only screen and (min-width : 768px) {}
		@media only screen and (min-width : 992px) {}
		@media only screen and (min-width : 1200px) {}
		</style>
  </body>
</html>
`

func main() {
	jsonFile, err := ioutil.ReadFile("gopher.json")

	j, err := parser.JSONParser(jsonFile)
	if err != nil {
		log.Fatalf("Failed to open JSON file \n %v", err)
	}

	h := web.Handler(j, tmpl)
	s := web.Server(h)

	log.Printf("Starting server on %s", s.Addr)
	log.Fatal(s.ListenAndServe())
}
