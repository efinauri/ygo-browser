package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./cards.cdb")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
