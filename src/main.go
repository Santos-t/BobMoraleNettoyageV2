package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	Email       string
	Client      bool
}

type Building struct {
	Address    string
	Complement string
	FloorNb    int
	Owner      User
}

type Ticket struct {
	Owner       User
	Address     Building
	Img         string
	Floor       int
	Status      string
	Orientation string
	Date        string
}

var path = "src/templates"
var home = template.Must(template.ParseFiles(path + "/home.html"))
var buildings = template.Must(template.ParseFiles(path + "/buildings.html"))
var profile = template.Must(template.ParseFiles(path + "/profile.html"))
var myTickets = template.Must(template.ParseFiles(path + "/myTickets.html"))
var planning = template.Must(template.ParseFiles(path + "/planning.html"))
var ticketForm = template.Must(template.ParseFiles(path + "/ticketForm.html"))

//var login = template.Must(template.ParseFiles(path + "/login.html"))
//var signup = template.Must(template.ParseFiles(path + "/signup.html"))
//var submittedTickets = template.Must(template.ParseFiles(path + "/submittedTickets.html"))

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/buildings", buildingsHandler)
	r.HandleFunc("/profile", profileHandler)
	r.HandleFunc("/myTickets", myTicketsHandler)
	r.HandleFunc("/planning", planningHandler)
	r.HandleFunc("/ticketForm", ticketFormHandler)

	http.ListenAndServe(":80", r)
}
