package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "1234"

func main() {
	port_formatted := fmt.Sprintf("0.0.0.0:%s", PORT)
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
	err = http.ListenAndServe(port_formatted, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Le serveur vient de recevoir une requete")
		fmt.Fprintln(w, "Le serveur renvoie la reponse suivante")
		count := store.AppendDB()
		fmt.Fprintf(w, "The database has %d values stored", count)
	}))
	if err != nil {
		log.Fatal("Cannot open the server - ", err)
	}
	fmt.Println("Last line of the file, does the program goes through this ?")
}
