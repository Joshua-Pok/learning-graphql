package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // use _ import for drivers
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Error initializing database: %v", err)
	}

	if db != nil {
		defer db.Close()
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Database not reachable: %v", err)
	}

}
