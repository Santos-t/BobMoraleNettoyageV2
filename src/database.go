package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Lien tuto : v
//https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/

func initDb() {
	//creating and opening a local database called mydb.db using the sqlite3 driver for Go
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	//creating the tables of the db
	//Table client
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS client(" +
		"id INTEGER PRIMARY KEY, " +
		"first_name TEXT NOT NULL, " +
		"last_name TEXT NOT NULL, " +
		"phone_number TEXT NOT NULL, " +
		"address TEXT, " +
		"role TEXT " +
		")",
	)
	checkErr(err)
	statement.Exec()

	//Table building
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS building (" +
		"id INTEGER PRIMARY KEY, " +
		"address TEXT NOT NULL, " +
		"complement TEXT, " +
		"floor_nb INTEGER, " +
		"owner_id INTEGER " +
		")",
	)
	checkErr(err)
	statement.Exec()

	//Table ticket
	statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS ticket(" +
		"id INTEGER PRIMARY KEY, " +
		"client_id INTEGER, " +
		"building_id INTEGER, " +
		"floor INTEGER, " +
		"img_path TEXT)",
	)
	checkErr(err)
	statement.Exec()
}

func insertClient(firstName, lastName, phoneNumber, address, role string) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO client(" +
		"firstName, lastName, phoneNumber, address, role)" +
		"VALUES(?, ?, ?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(firstName, lastName, phoneNumber, address, role)
	getClient()
}

func insertBuilding(address, complement string, floorNb, ownerId int) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO building" +
		"(address, complement, floorNb, ownerId) VALUES" +
		"(?, ?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(address, complement, floorNb, ownerId)
}

func insertTicket(clientId, buildingId, floor int) {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	statement, err := database.Prepare("INSERT INTO ticket" +
		"(clientId, buildingId, floor) VALUES" +
		"(?, ?, ?)",
	)
	checkErr(err)
	statement.Exec(clientId, buildingId, floor)
}

func getClient() {
	database, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)

	rows, err := database.Query("SELECT * FROM client")
	checkErr(err)

	var id int
	var first_name string
	var last_name string
	var phone_number string
	var address string
	var role string
	for rows.Next() {
		err = rows.Scan(&id, &first_name, &last_name, &phone_number, &address, &role)
		fmt.Println(strconv.Itoa(id) + ": " + first_name + " " + last_name + " " + phone_number + " " + address + " " + role)
	}
	rows.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*func main() {
  initDb()

  // Comment ajouter un client ==>

  first_name := "Bob"
  last_name := "Moral"
  phone_number := "0635284956"
  address := "2 rue des petits cailloux"
  role := "admin"
  insert_client(first_name, last_name, phone_number, address, role)
}*/
