package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Title string
}

var path = "src/templates"
var home = template.Must(template.ParseFiles(path + "/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Title: "coucou",
	}
	home.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	http.ListenAndServe(":80", r)
}
