package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"path/filepath"
	"yugioh-browser/database"
	"yugioh-browser/handlers"
)

const PORT = ":8088"

// Custom file server to set correct MIME type for CSS files
func fileServerWithMIMEType(root http.FileSystem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}
		http.FileServer(root).ServeHTTP(w, r)
	}
}

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()
	log.Println("Server is running: http://localhost" + PORT)

	router := chi.NewMux()
	router.Get("/", handlers.IndexHandler)
	router.Get("/api/cards", handlers.CardHandler(db))
	fs := http.Dir("static")
	router.Handle("/static/*", http.StripPrefix("/static/", fileServerWithMIMEType(fs)))

	log.Fatal(http.ListenAndServe(PORT, router))
}
