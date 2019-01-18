package main

import (
	"net/http"
)

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
		Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, Owner: current,
	}
	var tickets = []Ticket{
		Ticket{
			Owner: current, Address: address, Img: "img", Floor: 3, Status: "En cours", Orientation: "NNE", Date: "2 mars",
		},
	}
	data := MyTicketsData{
		TicketList: tickets,
	}
	myTickets.Execute(w, data)
}

func planningHandler(w http.ResponseWriter, r *http.Request) {
	var current = User{
		ID: 1, FirstName: "Antoine", LastName: "Legrand", PhoneNumber: "06", Address: "3 rue Gazan", Client: true,
	}
	var address = Building{
		Address: "3 rue gazan", Complement: "Bat. C", FloorNb: 7, Owner: current,
	}
	var tickets = []Ticket{
		Ticket{
			Owner: current, Address: address, Img: "img", Floor: 3, Status: "En cours", Orientation: "NNE", Date: "2 mars",
		},
	}
	data := PlanningData{
		TicketList: tickets,
	}
	planning.Execute(w, data)
}

func ticketFormHandler(w http.ResponseWriter, r *http.Request) {
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
	ticketForm.Execute(w, data)
}
