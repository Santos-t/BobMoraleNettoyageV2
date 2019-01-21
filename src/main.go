package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var path = "src/templates"
var home = template.Must(template.ParseFiles(path + "/home.html"))
var buildings = template.Must(template.ParseFiles(path + "/buildings.html"))
var profile = template.Must(template.ParseFiles(path + "/profile.html"))
var myTickets = template.Must(template.ParseFiles(path + "/myTickets.html"))
var planning = template.Must(template.ParseFiles(path + "/planning.html"))
var ticketForm = template.Must(template.ParseFiles(path + "/ticketForm.html"))
var submittedTickets = template.Must(template.ParseFiles(path + "/submittedTickets.html"))
var login = template.Must(template.ParseFiles(path + "/login.html"))
var signup = template.Must(template.ParseFiles(path + "/signup.html"))

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/", basicAuth(homeHandler))
	r.HandleFunc("/buildings", buildingsHandler)
	r.HandleFunc("/profile", profileHandler)
	r.HandleFunc("/myTickets", myTicketsHandler)
	r.HandleFunc("/planning", planningHandler)
	r.HandleFunc("/ticketForm", ticketFormHandler)
	r.HandleFunc("/submittedTickets", submittedTicketsHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/signup", signupHandler)

	http.ListenAndServe(":80", r)
}
