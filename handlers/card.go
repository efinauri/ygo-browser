package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"yugioh-browser/models"
)

func GetCardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
			return
		}

		card, err := models.GetCardByID(db, id)
		if err != nil {
			http.Error(w, "Error retrieving card", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.New("card").Parse(`
            <div class="card">
                <div class="card-name">{{.Name}}</div>
                <div class="card-desc">{{.Desc}}</div>
                <div class="card-types">Types: {{range .Types}}{{.}} {{end}}</div>
            </div>
        `)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, card)
	}
}
