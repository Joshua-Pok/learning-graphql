package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "blog.db")
	if err != nil {
		log.Fatal(err)
	}

	schema, err := os.ReadFile("./migrations/001_init.sql")
	if err != nil {
		log.Fatal(err)
	}

	db.MustExec(string(schema))

	migration, err := os.ReadFile("./migrations/002_comments.sql")
	if err != nil {
		log.Fatalf(err)
	}

	db.MustExec(string(migration))
	return db
}
