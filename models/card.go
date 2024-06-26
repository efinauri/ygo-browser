package models

import "database/sql"

type Card struct {
	ID    int      `json:"id"`
	OT    string   `json:"ot"`
	Alias string   `json:"alias"`
	Name  string   `json:"name"`
	Desc  string   `json:"desc"`
	Types []string `json:"types"`
}

func GetCardByID(db *sql.DB, id int) (*Card, error) {
	query := `
        SELECT d.id, d.ot, d.alias, d.type, t.name, t.desc
        FROM datas d
        LEFT JOIN texts t ON d.id = t.id
        WHERE d.id = ?`
	row := db.QueryRow(query, id)

	var card Card
	var typeMask int
	err := row.Scan(&card.ID, &card.OT, &card.Alias, &typeMask, &card.Name, &card.Desc)
	if err != nil {
		return nil, err
	}

	card.Types = DecodeCardTypes(typeMask)
	return &card, nil
}
