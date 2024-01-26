package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	// TODO: Change that when you are done
	connStr := ":memory:"
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		log.Fatal("Cannot connect to the database - ", err)
	}
	return &Store{db}, nil
}

func (s *Store) Ping() {
	if err := s.db.Ping(); err != nil {
		log.Fatal("Cannot ping to the database - ", err)
	}
}

func (s *Store) Init() {
	query := "CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, name TEXT NOT NULL);"
	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatal("Fail executing the query - ", err)
	}
}

func (s *Store) FillDB() {
	queries := []string{
		"INSERT INTO users (name) VALUES ('first query');",
		"INSERT INTO users (name) VALUES ('second query');",
		"INSERT INTO users (name) VALUES ('third query');",
		"INSERT INTO users (name) VALUES ('fourth query');",
	}
	for _, query := range queries {
		_, err := s.db.Exec(query)
		if err != nil {
			log.Fatal("Fail executing the query - ", err)
		}
	}
}

func (s *Store) CheckDB(value int) (id int, res string) {
	newquery := fmt.Sprintf("select * from users where id = %d", value)
	if err := s.db.QueryRow(newquery).Scan(&id, &res); err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No row found for ID %d", value)
		}
		log.Fatal(err)
	}
	return id, res
}
