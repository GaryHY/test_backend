package main

import (
	"fmt"
	"log"
)

func main() {
	// TODO: Mettre ce truc en production avec la base de donnee
	store, err := NewStore()
	defer store.db.Close()
	if err != nil {
		log.Fatal("Failed to create the new store - ", err)
	}
	store.Ping()
	store.Init()
	store.FillDB()
	id, queryRes := store.CheckDB(1)
	fmt.Printf("Value of the column with the value %d : %s\n", id, queryRes)
}
