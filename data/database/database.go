package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *gorm.DB
	err error
)

func OpenDb() (*gorm.DB, error) {
	db, err = gorm.Open("sqlite3", "inventory.db")
	db.LogMode(true)
	return db, err
}

func CloseDb() {
	db.Close()
}
