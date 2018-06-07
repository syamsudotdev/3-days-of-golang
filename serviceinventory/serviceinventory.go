package serviceinventory

import (
	"database/sql"
	"ijahinventory/database"
	"ijahinventory/model/log/ingoing"
	"ijahinventory/model/log/outgoing"
	"ijahinventory/model/product"
)

var (
	db  *sql.DB
	err error
)

func openDb() error {
	db, err = database.OpenDb()

	return err
}

func closeDb() {
	db.Close()
}
