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
		log.Fatal("could not connect to the Database", err)
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
		fmt.Println("ID: %d, Title: %s, Content: %s\n, created_at: %s, updated_at :%s", id, title, content, &created_at, &updated_at)
	}
}
