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

//var login = template.Must(template.ParseFiles(path + "/login.html"))
//var signup = template.Must(template.ParseFiles(path + "/signup.html"))
//var planning = template.Must(template.ParseFiles(path + "/planning.html"))
//var submittedTickets = template.Must(template.ParseFiles(path + "/submittedTickets.html"))
//var ticketForm = template.Must(template.ParseFiles(path + "/ticketForm.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		IsAClient: true,
	}
	home.Execute(w, data)
}

func buildingsHandler(w http.ResponseWriter, r *http.Request) {
	var owner = User{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var list = []Building{
		Building{
			Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, Owner: owner,
		},
	}
	data := BuildingsData{
		BuildingList: list,
	}
	buildings.Execute(w, data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	var current = User{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var list = []Building{
		Building{
			Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, Owner: owner,
		},
	}
	data := ProfileData{
		CurrentUser: current,
	}
	profile.Execute(w, data)
}

func myTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var current = User{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var address = Building{
		Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, Owner: owner,
	}
	var ticket = Ticket{
		Owner: current, Address: address, Img: "img", Floor: 3, Status: "En cours", Orientation: "NNE", Date: "2 mars" 
	}
	data := MyTicketsData{
		TicketList: tickets,
	}
	myTickets.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/buildings", buildingsHandler)
	r.HandleFunc("/profile", profileHandler)

	http.ListenAndServe(":80", r)
}
