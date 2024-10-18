package main

import (
	"log"
	"notes-app/db"
)

func main() {

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatal("could not connect to the Database\n", err)
	}
	defer dbConn.Close()

	rows, err := dbConn.Query("SELECT id, title, content, created_at, updated_at FROM notes")

	if err != nil {
		log.Fatal("Nothing is present in the DB", err)

	}
	defer rows.Close()
}
