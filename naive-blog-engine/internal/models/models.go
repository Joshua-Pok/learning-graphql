package models

type Post struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	Created_at string `db:"created_at"`
}

type Comment struct {
	Id      int    `db:"id"`
	post_id int    `db:"post_id"`
	Content string `db:"content"`
}
