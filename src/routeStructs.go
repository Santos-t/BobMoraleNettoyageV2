package main

type HomeData struct {
	IsAClient bool
}

type BuildingsData struct {
	BuildingList []Building
}

type ProfileData struct {
	CurrentUser User
}

type MyTicketsData struct {
	TicketList []Ticket
}

type PlanningData struct {
	TicketList []Ticket
}

type TicketFormData struct {
	BuildingList []Building
}

type submittedTicketsData struct {
	TicketList []Ticket
}

type loginData struct {
}

type signupData struct {
}
