package gql

import (
	"github.com/Joshua-Pok/naive-blog/internal/models"
	"github.com/graphql-go/graphql"
	"github.com/jmoiron/sqlx"
)

var postType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func NewSchema(database *sqlx.DB) (graphql.Schema, error) {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"post": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)

					var post models.Post

					err := database.Get(&post, "SELECT * FROM posts WHERE id = ?", id)
					if err != nil {
						return nil, err
					}

					return post, nil
				},
			},
		},
	})
	cfg := graphql.SchemaConfig{
		Query: rootQuery,
	}

	return graphql.NewSchema(cfg)
}
