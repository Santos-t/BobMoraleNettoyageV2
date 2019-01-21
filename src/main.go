package main

import (
	"encoding/base64"
	"html/template"
	"net/http"
	"strings"

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
	ID         int
	Address    string
	Complement string
	FloorNb    int
	Owner      User
}

type Ticket struct {
	ID          int
	Owner       User
	Address     Building
	Img         string
	Floor       int
	Status      string
	Orientation string
	Date        string
}

func basicAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Not authorized", 401)
			return
		}

		if pair[0] != "username" || pair[1] != "password" {
			http.Error(w, "Not authorized", 401)
			return
		}

		h.ServeHTTP(w, r)
	}
}

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
