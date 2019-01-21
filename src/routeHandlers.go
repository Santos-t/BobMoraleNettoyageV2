package main

import (
	"net/http"
	"strconv"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		IsAClient: true,
	}
	home.Execute(w, data)
}

func buildingsHandler(w http.ResponseWriter, r *http.Request) {
	var owner = Client{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var list = []Building{
		Building{
			ID: 1, Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, ClientId: owner.ID,
		},
	}
	data := BuildingsData{
		BuildingList: list,
	}
	buildings.Execute(w, data)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	var current = Client{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	data := ProfileData{
		CurrentUser: current,
	}
	profile.Execute(w, data)
}

func myTicketsHandler(w http.ResponseWriter, r *http.Request) {
	list := getTicketsFromUser(idUser)
	data := MyTicketsData{
		TicketList: list,
	}
	myTickets.Execute(w, data)
}

func planningHandler(w http.ResponseWriter, r *http.Request) {
	var current = Client{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var address = Building{
		Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, ClientId: current.ID,
	}
	var tickets = []Ticket{
		Ticket{
			OwnerId: current.ID, BuildingId: address.ID, Img: "img", Floor: 3, Status: "En cours", Orientation: "NNE", Date: time.Now(),
		},
	}
	data := PlanningData{
		TicketList: tickets,
	}
	planning.Execute(w, data)
}

func ticketFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var owner = Client{
			ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
		}
		var list = []Building{
			Building{
				ID: 1, Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, ClientId: owner.ID,
			},
		}
		data := BuildingsData{
			BuildingList: list,
		}
		ticketForm.Execute(w, data)
	} else {
		bu := r.FormValue("building")
		fl := r.FormValue("floor")
		orientation := r.FormValue("orientation")
		building, err := strconv.Atoi(bu)
		floor, err := strconv.Atoi(fl)
		if err != nil {
			print("Form not conform")
		}
		var ticket = Ticket{
			OwnerId: idUser, BuildingId: building, Img: "img", Floor: floor, Status: "En cours", Orientation: orientation, Date: time.Now(),
		}
		insertTicket(ticket)

		data := HomeData{
			IsAClient: true,
		}
		home.Execute(w, data)
	}
}

func submittedTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var current = Client{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var address = Building{
		Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, ClientId: current.ID,
	}
	var tickets = []Ticket{
		Ticket{
			OwnerId: current.ID, BuildingId: address.ID, Img: "img", Floor: 3, Status: "En cours", Orientation: "NNE", Date: time.Now(),
		},
	}
	data := PlanningData{
		TicketList: tickets,
	}
	submittedTickets.Execute(w, data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	data := loginData{}
	login.Execute(w, data)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	data := signupData{}
	signup.Execute(w, data)
}
