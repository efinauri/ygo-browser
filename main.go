package main

import (
	"log"
	"net/http"
	"yugioh-browser/database"
	"yugioh-browser/handlers"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/card", handlers.GetCardHandler(db))

	log.Println("Server is running on port 8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
