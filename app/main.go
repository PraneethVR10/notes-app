package main

import (
	"fmt"
	"log"
	"notes-app/db"
	"time"
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

	for rows.Next() {
		var id int
		var title, content string
		var created_at, updated_at time.Time

		err := rows.Scan(&id, &title, &content, &created_at, &updated_at)
		if err != nil {
			log.Fatal("error scanning row", err)

		}
		fmt.Printf(" ID: %d,\n Title: %s,\n Content: %s,\n created_at: %s,\n updated_at :%s\n", id, title, content, &created_at, &updated_at)
	}
}
