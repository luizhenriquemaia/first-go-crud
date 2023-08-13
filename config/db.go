package config

import (
	"database/sql"
	"log"
)

const filename = "sqlite.db"

func GetDB() *sql.DB {

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitDB() {
	db := GetDB()
	repository := NewSqliteRepository(db)

	if err := repository.Migrate(); err != nil {
		log.Fatal(err)
	}
}
