package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "8080"

func main() {
	port_formatted := fmt.Sprintf(":%s", PORT)
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
	// TODO: On teste d'abord le serveur et ensuite on gere la connection avec la base de donnee
	if err := http.ListenAndServe(port_formatted, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Je tape dans le serveur")
		fmt.Fprintln(w, "Le serveur renvoit la reponse suivante")
	})); err != nil {
		log.Fatal("Cannot connect to the server - %s", err)
	}
}
