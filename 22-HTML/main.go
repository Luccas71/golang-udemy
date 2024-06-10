package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuario{
			Nome:  "Lucas",
			Email: "lucas@email.com",
		}
		templates.ExecuteTemplate(w, "index.html", u)
	})
	fmt.Println("Rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
