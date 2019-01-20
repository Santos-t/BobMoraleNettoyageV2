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
  database, err := sql.Open("sqlite3", "./mydb.db")
  check_err(err)

  //creating the tables of the db
  //Table client
  statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS client(" +
    "id INTEGER PRIMARY KEY, " +
    "first_name TEXT NOT NULL, " +
    "last_name TEXT NOT NULL, " +
    "phone_number TEXT NOT NULL, " +
    "address TEXT, " +
    "role TEXT " +
    ")")
    check_err(err)
    statement.Exec()

    //Table building
    statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS building (" +
      "id INTEGER PRIMARY KEY, " +
      "address TEXT NOT NULL, " +
      "complement TEXT, " +
      "floor_nb INTEGER, " +
      "owner_id INTEGER " +
      ")")
      check_err(err)
      statement.Exec()

      //Table ticket
      statement, err = database.Prepare("CREATE TABLE IF NOT EXISTS ticket(" +
        "id INTEGER PRIMARY KEY, " +
        "client_id INTEGER, " +
        "building_id INTEGER, " +
        "floor INTEGER, " +
        "img_path TEXT)")
        check_err(err)
        statement.Exec()
      }

      func insert_client(first_name, last_name, phone_number, address, role string){
        database, err := sql.Open("sqlite3", "./mydb.db")
        check_err(err)

        statement, err := database.Prepare("INSERT INTO client(" +
          "first_name, last_name, phone_number, address, role)" +
          "VALUES(?, ?, ?, ?, ?)")
        check_err(err)
        statement.Exec(first_name, last_name, phone_number, address, role)
        get_client()
      }

      func insert_building(address, complement string, floor_nb, owner_id int){
        database, err := sql.Open("sqlite3", "./mydb.db")
        check_err(err)

        statement, err := database.Prepare("INSERT INTO building" +
          "(address, complement, floor_nb, owner_id) VALUES" +
          "(?, ?, ?, ?)")
        check_err(err)
        statement.Exec(address, complement, floor_nb, owner_id)
      }

      func insert_ticket(client_id, building_id, floor int){
        database, err := sql.Open("sqlite3", "./mydb.db")
        check_err(err)

        statement, err := database.Prepare("INSERT INTO ticket" +
          "(client_id, building_id, floor) VALUES" +
          "(?, ?, ?)")
        check_err(err)
        statement.Exec(client_id, building_id, floor)
      }

      func get_client() {
        database, err := sql.Open("sqlite3", "./mydb.db")
        check_err(err)

        rows, err := database.Query("SELECT * FROM client")
        check_err(err)

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

      func check_err(err error) {
        if err != nil {
          panic(err)
        }
      }

      func main() {
        init_db()

        first_name := "Bob"
        last_name := "Moral"
        phone_number := "0635284956"
        address := "2 rue des petits cailloux"
        role := "admin"
        insert_client(first_name, last_name, phone_number, address, role)

      }
