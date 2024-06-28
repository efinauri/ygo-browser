package handlers

import (
	"net/http"
	views "yugioh-browser/views"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	views.Index().Render(r.Context(), w)
}
