package services

import (
	"database/sql"
	"log"
	"log/slog"
	"strings"
	"yugioh-browser/models/dtos"
	"yugioh-browser/models/endpoint_params"
	"yugioh-browser/models/entities"
)

const SelectCards = `
SELECT d.id, d.type, t.name, t.desc, d.atk, d.def, d.level, d.race, d.attribute
FROM datas d
LEFT JOIN texts t ON d.id = t.id 
`

const CountCards = `
SELECT COUNT(*)
FROM datas d
LEFT JOIN texts t ON d.id = t.id
`

func MapRowsToCards(rows *sql.Rows) []*entities.Card {
	var cards []*entities.Card
	for rows.Next() {
		var card entities.Card
		var typeMask int
		var raceMask int
		var attributeMask int
		err := rows.Scan(&card.ID, &typeMask, &card.Name, &card.Desc, &card.Atk, &card.Def, &card.Level, &raceMask, &attributeMask)
		if err != nil {
			return cards
		}
		card.Sanitize(typeMask, raceMask, attributeMask)
		cards = append(cards, &card)
	}
	return cards
}

func GetAllCards(db *sql.DB, filters endpoint_params.CardSearchFilters, page int, pageSize int) dtos.PaginatedCardResult {
	offset := (page - 1) * pageSize
	var queryFragments []string
	var args []interface{}

	if filters.Name != "" {
		queryFragments = append(queryFragments, "t.name like ?")
		args = append(args, "%"+filters.Name+"%")
	}
	if filters.AtkGt > 0 {
		queryFragments = append(queryFragments, "d.atk >= ?")
		args = append(args, filters.AtkGt)
	}
	if filters.AtkLte > 0 {
		queryFragments = append(queryFragments, "d.atk <= ?")
		args = append(args, filters.AtkLte)
	}
	if filters.DefGt > 0 {
		queryFragments = append(queryFragments, "d.def >= ?")
		args = append(args, filters.DefGt)
	}
	if filters.DefLte > 0 {
		queryFragments = append(queryFragments, "d.def <= ?")
		args = append(args, filters.DefLte)
	}
	if filters.LvGt > 0 {
		queryFragments = append(queryFragments, "d.level >= ?")
		args = append(args, filters.LvGt)
	}
	if filters.LvLte > 0 {
		queryFragments = append(queryFragments, "d.level <= ?")
		args = append(args, filters.LvLte)
	}

	selectQuery := SelectCards
	countQuery := CountCards
	if len(queryFragments) > 0 {
		queryExtension := "WHERE " + strings.Join(queryFragments, "AND ")
		selectQuery += queryExtension
		countQuery += queryExtension
	}

	totalElements := 0
	slog.Info(countQuery)
	err := db.QueryRow(countQuery, args...).Scan(&totalElements)
	if err != nil {
		log.Fatal(err)
	}

	selectQuery += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	slog.Info(selectQuery)
	rows, err := db.Query(selectQuery, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cards := MapRowsToCards(rows)
	return dtos.PaginatedCardResult{Elements: cards, Page: page, Size: pageSize, TotalElements: totalElements}
}
