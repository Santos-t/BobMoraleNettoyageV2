package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type HomeData struct {
	IsAClient bool
}

var path = "src/templates"
var home = template.Must(template.ParseFiles(path + "/home.html"))

//var login = template.Must(template.ParseFiles(path + "/login.html"))
//var signup = template.Must(template.ParseFiles(path + "/signup.html"))
//var buildings = template.Must(template.ParseFiles(path + "/buildings.html"))
//var myTickets = template.Must(template.ParseFiles(path + "/myTickets.html"))
//var planning = template.Must(template.ParseFiles(path + "/planning.html"))
//var profile = template.Must(template.ParseFiles(path + "/profile.html"))
//var submittedTickets = template.Must(template.ParseFiles(path + "/submittedTickets.html"))
//var ticketForm = template.Must(template.ParseFiles(path + "/ticketForm.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		IsAClient: true,
	}
	home.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	http.ListenAndServe(":80", r)
}
