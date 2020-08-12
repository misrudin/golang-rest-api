package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/node_api")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
