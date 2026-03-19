package schema

import (
	"fmt"

	"github.com/Joshua-Pok/library-api/internal/data"
	"github.com/Joshua-Pok/library-api/internal/models"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{ //object is basically a typek
	Name: "RootQuery",
	Fields: graphql.Fields{
		"hello": &graphql.Field{ //field is a piece of data we can ask for
			Type:    graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return "World", nil }, //Resolve is a function assigned to a field that is executed when that field is queried
		},

		"book": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				for _, book := range data.Books {
					if book.ID == id {
						book.AuthorID = "1"
						return book, nil
					}
				}
				return nil, nil

			},
		},
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return data.Books, nil
			},
		},
	},
})

var authorType = graphql.NewObject(graphql.ObjectConfig{
	Name: "authorType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "bookType",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: authorType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				token := p.Context.Value(userkey)
				fmt.Printf("TYPE: %T\n", p.Source)
				book, ok := p.Source.(models.Book)
				if !ok {
					fmt.Println("Not a book unable to cast")
				}
				fmt.Printf("Book author ID: %s\n", book.AuthorID)
				for _, author := range data.Authors {
					if author.ID == book.AuthorID {
						return author, nil
					}
				}

				return nil, nil

			},
		},
	},
})

func NewSchema() (graphql.Schema, error) {
	cfg := graphql.SchemaConfig{
		Query: rootQuery,
	}
	return graphql.NewSchema(cfg)

}
