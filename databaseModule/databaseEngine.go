package databaseModule

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Db *sql.DB

func init() {
	// open database file
	var err error
	Db, err = sql.Open("sqlite3", "./resources/database.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer Db.Close()
}



func main() {
	var err error
	defer Db.Close()
	// create script TODO: add file.sql
	sqlStmt := `CREATE TABLE data (
  		id integer PRIMARY KEY AUTOINCREMENT,
  		did integer,
  		date datetime,
  		data varchar,
  		data_type varchar
	);`
	// exec script
	_, err = Db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// create a new one table
	sqlStmt = `CREATE TABLE device (
		id integer PRIMARY KEY AUTOINCREMENT,
		device_name varchar,
		device_condition varchar
	);`
	// exec script
	_, err = Db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

}
