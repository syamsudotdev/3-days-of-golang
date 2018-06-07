package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"ijahinventory/model/log/ingoing"
	"ijahinventory/model/log/outgoing"
	"ijahinventory/model/product"
)

var (
	db  *gorm.DB
	err error
)

func OpenDb() (*gorm.DB, error) {
	db, err = gorm.Open("sqlite3", "ijah-inventory.db")
	db.AutoMigrate()
	db.LogMode(true)
	return db, err
}

func CloseDb() {
	db.Close()
}
