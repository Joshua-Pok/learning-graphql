package schema

import "github.com/graphql-go/graphql"
import "github.com/Joshua-Pok/library-api/internal/models"

var books = []models.Book{
	{ID: "1", Title: "1984"},
	{ID: "2", Title: "The Hobbit"},
}

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
				for _, book := range books {
					if book.ID == id {
						return book, nil
					}
				}
				return nil, nil

			},
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
	},
})

func NewSchema() (graphql.Schema, error) {
	cfg := graphql.SchemaConfig{
		Query: rootQuery,
	}
	return graphql.NewSchema(cfg)

}
