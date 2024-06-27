package models

import (
	"database/sql"
	"log"
	"strings"
	"yugioh-browser/models/enums"
)

type Card struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Desc  string   `json:"desc"`
	Types []string `json:"types"`
	Atk   int      `json:"atk"`
	Def   int      `json:"def"`
	Level int      `json:"level"`
}

type PagedCards struct {
	Elements []*Card
	Page     int
	Size     int
}

type CardFilters struct {
	AtkGt  int
	AtkLte int
	DefGt  int
	DefLte int
	LvGt   int
	LvLte  int
	Name   string
}

const SelectCard = `
SELECT d.id, d.type, t.name, t.desc, d.atk, d.def, d.level
FROM datas d
LEFT JOIN texts t ON d.id = t.id 
`

func (c Card) Sanitize(typeMask int) {
	c.Types = enums.DecodeCardTypes(typeMask)
}

func MapRowsToCards(rows *sql.Rows) []*Card {
	var cards []*Card
	for rows.Next() {
		var card Card
		var typeMask int
		err := rows.Scan(&card.ID, &typeMask, &card.Name, &card.Desc, &card.Atk, &card.Def, &card.Level)
		if err != nil {
			return cards
		}
		card.Sanitize(typeMask)
		cards = append(cards, &card)
	}
	return cards
}

func GetCardsByFilters(db *sql.DB, page int, pageSize int, filters CardFilters) PagedCards {
	offset := (page - 1) * pageSize
	query := ""
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

	if len(queryFragments) > 0 {
		query = SelectCard + "WHERE " + strings.Join(queryFragments, "AND ")
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	log.Println("running the following query:\n", query)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	cards := MapRowsToCards(rows)
	return PagedCards{Elements: cards, Page: page, Size: pageSize}
}
