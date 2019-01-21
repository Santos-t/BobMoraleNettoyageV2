package main

import (
	"database/sql"
	"container/list"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Lien tuto : v
//https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/

type Client struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
	Address     string
	Email       string
	Client      bool
	Password    string
}

type Building struct {
	ID 				 int
	Address    string
	Complement string
	FloorNb    int
	ClientId	 int
}

type Ticket struct {
	ID          int
	OwnerId     int
	BuildingId  int
	Img         string
	Floor       int
	Status      string
	Orientation string
	Date        time.Time
}

func initDb() {
	//creating and opening a local database called mydb.db using the sqlite3 driver for Go
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	//creating the tables of the db
	//Table client
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS client(" +
		"id INTEGER PRIMARY KEY, " +
		"firstName TEXT NOT NULL, " +
		"lastName TEXT NOT NULL, " +
		"phoneNumber TEXT NOT NULL, " +
		"address TEXT, " +
		"password TEXT, " +
		"email TEXT, " +
		"client BOOLEAN" +
		")",
	)
	checkErr(err)
	statement.Exec()

	//Table building
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS building (" +
		"id INTEGER PRIMARY KEY, " +
		"address TEXT NOT NULL, " +
		"complement TEXT, " +
		"floorNb INTEGER, " +
		"ownerId INTEGER" +
		")",
	)
	checkErr(err)
	statement.Exec()

	//Table ticket
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS ticket(" +
		"id INTEGER PRIMARY KEY, " +
		"clientId INTEGER, " +
		"buildingId INTEGER, " +
		"floor INTEGER, " +
		"orientation TEXT, " +
		"date DATETIME, " +
		"imgPath TEXT)",
	)
	checkErr(err)
	statement.Exec()
}

func insertClient(client Client) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO client(" +
		"firstName, lastName, phoneNumber, address, client)" +
		"VALUES(?, ?, ?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(client.FirstName, client.LastName, client.PhoneNumber, client.Address, client.Client)
}

func insertBuilding(building Building) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO building" +
		"(address, complement, floorNb, ownerId) VALUES" +
		"(?, ?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(building.Address, building.Complement, building.FloorNb, building.ClientId)
}

func insertTicket(ticket Ticket) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO ticket" +
		"(clientId, buildingId, floor, orientation, date, imgPath) VALUES" +
		"(?, ?, ?, ?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(ticket.OwnerId, ticket.BuildingId, ticket.Floor, ticket.Orientation, ticket.Date, ticket.Img)
}

func getClient() *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM client")
	checkErr(err)

	var id int
	var firstName string
	var lastName string
	var phoneNumber string
	var address string
	var role bool
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &phoneNumber, &address, &role)
		var owner = Client{
			ID: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Address: address, Client: role,
		}
		result.PushBack(owner)
	}
	rows.Close()
	return result
}

func getClientFromId(clientId int) *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM client WHERE id IS " + strconv.Itoa(clientId))
	checkErr(err)

	var id int
	var firstName string
	var lastName string
	var phoneNumber string
	var address string
	var role bool
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &firstName, &lastName, &phoneNumber, &address, &role)
		var owner = Client{
			ID: id, FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Address: address, Client: role,
		}
		result.PushBack(owner)
	}
	rows.Close()
	return result
}

//building d'un username
func getBuildingFromUser(userId int) *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM building WHERE ownerId IS " + strconv.Itoa(userId))
	checkErr(err)

	var id int
	var address string
	var complement string
	var floorNb int
	var clientId int
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &address, &complement, &floorNb, &clientId)
		var building = Building{
			ID: id, Address: address, Complement: complement, FloorNb: floorNb, ClientId: clientId,
		}
		result.PushBack(building)
	}
	rows.Close()
	return result
}

func getBuildingFromId(buildingId int) *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM building WHERE is IS " + strconv.Itoa(buildingId))
	checkErr(err)

	var id int
	var address string
	var complement string
	var floorNb int
	var clientId int
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &address, &complement, &floorNb, &clientId)
		var building = Building{
			ID: id, Address: address, Complement: complement, FloorNb: floorNb, ClientId: clientId,
		}
		result.PushBack(building)
	}
	rows.Close()
	return result
}

//ticket all et d'un user
func getTickets() *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM ticket")
	checkErr(err)

	var id int
	var clientId int
	var buildingId int
	var floor int
	var orientation string
	var date time.Time
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &clientId, &buildingId, &floor, &orientation, &date)
		var ticket = Ticket{
			ID: id, OwnerId: clientId, BuildingId: buildingId, Floor: floor, Orientation: orientation, Date: date,
		}
		result.PushBack(ticket)
	}
	rows.Close()
	return result
}

func getTicketsFromUser(userId int) *list.List{
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM ticket WHERE clientId IS " + strconv.Itoa(userId))
	checkErr(err)

	var id int
	var clientId int
	var buildingId int
	var floor int
	var orientation string
	var date time.Time
	result := list.New()
	for rows.Next() {
		err = rows.Scan(&id, &clientId, &buildingId, &floor, &orientation, &date)
		var ticket = Ticket{
			ID: id, OwnerId: clientId, BuildingId: buildingId, Floor: floor, Orientation: orientation, Date: date,
		}
		result.PushBack(ticket)
	}
	rows.Close()
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
  initDb()
  // Comment ajouter un client ==>
  firstName := "Bob"
  lastName := "Moral"
  phoneNumber := "0635284956"
  address := "2 rue des petits cailloux"
  client := true
	var c = Client{
		ID: 0,  FirstName: firstName, LastName: lastName, PhoneNumber: phoneNumber, Address: address, Client: client,
	}
  insertClient(c)
	list := getClient()

	for e := list.Front(); e != nil; e = e.Next() {
		var client Client
		client = e.Value.(Client)
		fmt.Println(client.FirstName + " " + client.PhoneNumber)
	}
}
