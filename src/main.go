package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	title string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		title: "coucou",
	}
	tmpl := template.Must(template.ParseFiles("src/templates/index.html"))
	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.ListenAndServe(":80", r)
}
