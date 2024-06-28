package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"yugioh-browser/models/endpoint_params"
	"yugioh-browser/services"
	"yugioh-browser/views"
)

func CardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method == "GET" {
			page, pageSize := GetPaginationParams(r)
			filters := CollectCardFilters(r)
			HandleGetCardsByFilters(db, w, r, page, pageSize, filters)
		}
	}
}

func HandleGetCardsByFilters(db *sql.DB, w http.ResponseWriter, r *http.Request, page int, pageSize int, filters endpoint_params.CardSearchFilters) {
	cards := services.GetAllCards(db, filters, page, pageSize)
	views.CardsResult(cards).Render(r.Context(), w)
}

func GetPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	return page, pageSize
}

func CollectCardFilters(r *http.Request) endpoint_params.CardSearchFilters {
	filters := endpoint_params.CardSearchFilters{}
	if name := r.URL.Query().Get("name"); name != "" {
		filters.Name = name
	}
	if atkGt := r.URL.Query().Get("atk_gt"); atkGt != "" {
		filters.AtkGt, _ = strconv.Atoi(atkGt)
	}
	if atkLt := r.URL.Query().Get("atk_lt"); atkLt != "" {
		filters.AtkLte, _ = strconv.Atoi(atkLt)
	}
	if defGt := r.URL.Query().Get("def_gt"); defGt != "" {
		filters.DefGt, _ = strconv.Atoi(defGt)
	}
	if defLt := r.URL.Query().Get("def_lt"); defLt != "" {
		filters.DefLte, _ = strconv.Atoi(defLt)
	}
	if lvGt := r.URL.Query().Get("lv_gt"); lvGt != "" {
		filters.LvGt, _ = strconv.Atoi(lvGt)
	}
	if lvLt := r.URL.Query().Get("lv_lt"); lvLt != "" {
		filters.LvLte, _ = strconv.Atoi(lvLt)
	}
	return filters
}
