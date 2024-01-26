package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main_old() {
	connStr := ":memory:"
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		log.Fatal("Cannot connect to the database - ", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping to the database - ", err)
	}
	query := "CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL);"
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Fail executing the query - ", err)
	}
	query2 := "INSERT INTO users (name) VALUES ('first query');"
	query3 := "INSERT INTO users (name) VALUES ('second query');"
	query4 := "INSERT INTO users (name) VALUES ('third query');"
	query5 := "INSERT INTO users (name) VALUES ('fourth query');"
	queries := []string{query2, query3, query4, query5}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal("Fail executing the query - ", err)
		}
	}
}
