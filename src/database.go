package main

import (
  "database/sql"
  "fmt"
  "strconv"

  _ "github.com/mattn/go-sqlite3"
)

// Lien tuto : v
//https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/
func init_db() {
  //creating and opening a local database called mydb.db using the sqlite3 driver for Go
  database, _ := sql.Open("sqlite3", "./mydb.db")

  //creating the tables of the db
  //Table client
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS client (
    id INTEGER PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    phone_number TEXT,
    address TEXT,
    role TEXT,
    )")
  statement.Exec()

  //Table building
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS building (
    id INTEGER PRIMARY KEY,
    address TEXT,
    complement TEXT,
    floor_nb INTEGER,
    owner TEXT
    )")
  statement.Exec()

  //Table ticket
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS ticket (
    id INTEGER PRIMARY KEY,
    client TEXT,
    last_name TEXT,

    )")
  statement.Exec()


  //creating template of insertion that can be executer later
  statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")

  //execution of the template with values
  statement.Exec("Nic", "Raboy")
  rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
  var id int
  var firstname string
  var lastname string
  for rows.Next() {
      rows.Scan(&id, &firstname, &lastname)
      fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
}

func main() {
    init_db()
    }
}
