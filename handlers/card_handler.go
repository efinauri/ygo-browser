package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"yugioh-browser/models/dtos"
	"yugioh-browser/models/endpoint_params"
	"yugioh-browser/services"
)

func CardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			page, pageSize := GetPaginationParams(r)
			filters := CollectCardFilters(r)
			HandleGetCardsByFilters(db, w, page, pageSize, filters)
		}
	}
}

func HandleGetCardsByFilters(db *sql.DB, w http.ResponseWriter, page int, pageSize int, filters endpoint_params.CardSearchFilters) {
	cards := services.GetAllCards(db, filters, page, pageSize)
	UpdateDomWithResult(w, cards)
}

func UpdateDomWithResult(w http.ResponseWriter, cards dtos.PaginatedCardResult) {
	funcMap := template.FuncMap{
		"add": func(n int) int { return n + 1 },
		"sub": func(n int) int { return n - 1 },
	}
	tmpl, err := template.New("card").Funcs(funcMap).Parse(`
        {{ range .Elements }}
        <div class="card">
            <div class="card-name">{{.Name}}</div>
            {{ if gt .Level 0 }}
            <div class="card-level">Lv. {{.Level}} {{.Attribute}} {{.Race}}</div>
            {{ end }}
            <div class="card-types">{{.Types}}</div>
            <div class="card-desc">{{.Desc}}</div>
            {{ if and (gt .Atk 0) (gt .Def 0) }}
            <div class="card-stats">{{.Atk}}/{{.Def}}</div>
            {{ end }}
        </div>
        {{ end }}
        <div class="pagination">
            {{ if gt .Page 1 }}
            <button hx-get="/card?name={{$.Name}}&page={{sub .Page 1}}&pageSize={{.Size}}" hx-target="#result" hx-swap="innerHTML">Previous</button>
            {{ end }}
            <span>Page {{.Page}}</span>
            {{ if eq (len .Elements) .Size }}
            <button hx-get="/card?name={{$.Name}}&page={{add .Page 1}}&pageSize={{.Size}}" hx-target="#result" hx-swap="innerHTML">Next</button>
            {{ end }}
        </div>
    `)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_ = tmpl.Execute(w, cards)
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
