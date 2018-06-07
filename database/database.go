package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	err error
)

func OpenDb() (*sql.DB, error) {
	db, err = sql.Open("sqlite3", "ijah-inventory.db")
	return db, err
}

func CloseDb() {
	db.Close()
}
