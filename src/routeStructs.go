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
