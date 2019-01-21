package main

import "container/list"

type HomeData struct {
	IsAClient bool
}

type BuildingsData struct {
	BuildingList []Building
}

type ProfileData struct {
	CurrentUser Client
}

type MyTicketsData struct {
	TicketList *list.List
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
