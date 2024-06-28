package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"yugioh-browser/database"
	"yugioh-browser/handlers"
)

const PORT = ":8088"

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	log.Println("Server is running: http://localhost" + PORT)

	router := chi.NewMux()
	router.Get("/", handlers.IndexHandler)
	router.Get("/card", handlers.CardHandler(db))

	log.Fatal(http.ListenAndServe(PORT, router))
}
