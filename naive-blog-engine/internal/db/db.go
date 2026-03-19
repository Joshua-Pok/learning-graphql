package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	db, err := sqlx.Open("sqlite3", "blog.db")
	if err != nil {
		log.Fatal(err)
	}

	if db != nil {
		defer db.Close()
	}

	schema, err := os.ReadFile("./migrations/001_init.sql")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(string(schema))
	return db
}
