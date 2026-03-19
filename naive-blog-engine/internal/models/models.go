package models

type Post struct {
	id         int    `db:"id"`
	title      string `db:"title"`
	content    string `db:"content"`
	author     string `db:"author"`
	created_at string `db:"created_at"`
}
