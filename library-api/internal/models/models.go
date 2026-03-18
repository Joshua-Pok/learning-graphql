package models

type Book struct {
	ID       string
	Title    string
	AuthorID string
}

type Author struct {
	ID   string
	Name string
}
