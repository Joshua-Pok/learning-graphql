package main

import (
	"context"
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

func TestFetchAllPosts(t *testing.T) {
	database := db.InitDB()
	defer database.Close()

	database.MustExec(`
INSERT INTO posts (title, content, author, created_at) VALUES
(First Post", "This is the content of the first post", "Alice", "2026-03-19"),
(Second Post", "This is the content of the second post", "Bob", "2026-03-19"),
(Third Post", "This is the content of the third post", "Charlie", "2026-03-19");
`)

	database.MustExec(`
INSERT INTO comments (post_id, content) VALUES
(1", "Comment A1"),
(1", "Comment A2"),
(2", "Comment B1"),
(2", "Comment B2"),
(3", "Comment C1"),
(3", "Comment C2");
`)

	queryString := " { posts { title comments { content } } }"

	schema, err := gql.NewSchema(database)
	if err != nil {
		t.Errorf("Error Creating schema")
	}

	ctx := context.WithValue(context.Background(), "db", database)

	params := graphql.Params{
		Schema:        schema,
		RequestString: queryString,
		Context:       ctx,
	}

	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		t.Errorf("Error executing query: %v", res.Errors)

	}

	data := res.Data.(map[string]interface{})
	posts, ok := data["posts"].([]interface{}) //list of posts
	if !ok {
		t.Fatalf("Post filed missing or wrong type")
	}

	if len(posts) != 3 {
		t.Fatalf("expected 3 posts, got %d", len(posts))
	}

	for _, p := range posts {
		postMap := p.(map[string]interface{})

		comments, ok := postMap["comments"].([]interface{})
		if !ok {
			t.Fatalf("comments missing or wrong type")
		}

		if len(comments) != 2 {
			t.Fatalf("expected 2 comments, got %d", len(comments))

		}

	}
}
