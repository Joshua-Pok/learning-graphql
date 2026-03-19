package main

import (
	"database/sql"
	"testing"

	"github.com/Joshua-Pok/naive-blog/internal/db"
	"github.com/Joshua-Pok/naive-blog/internal/gql"
	"github.com/graphql-go/graphql"
	_ "github.com/mattn/go-sqlite3" // use _ import for drivers
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

func TestFetchPost(t *testing.T) {
	database := db.InitDB()
	defer database.Close()

	database.MustExec(`
        INSERT INTO posts (title, content, author)
        VALUES ('Test Title', 'Test Content', 'Author')
    `)
	queryString := "{ post (id: \"1\") { title content } }"

	schema, err := gql.NewSchema(database)
	if err != nil {
		t.Errorf("Error creating Schema: %v", err)
	}

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryString,
	}

	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		t.Errorf("Error executing query: %v", res.Errors)
	}

	data := res.Data.(map[string]interface{})
	post, ok := data["post"].(map[string]interface{})
	if !ok {
		t.Fatalf("Post field missing or wrong type")
	}

	if post["title"] == nil {
		t.Errorf("expected title, got nil")
	}

	if post["content"] == nil {
		t.Errorf("expected content, got nil")
	}

}
