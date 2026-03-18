package data

import "github.com/Joshua-Pok/library-api/internal/models"

var Authors = []models.Author{
	{
		ID:   "1",
		Name: "JKRowling",
	},
}

var Books = []models.Book{
	{
		ID:       "1",
		Title:    "The go programming language",
		AuthorID: "1",
	},
	{
		ID:       "2",
		Title:    "your mom gay",
		AuthorID: "1",
	},
}
