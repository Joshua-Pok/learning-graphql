package gql

import (
	"fmt"
	"log"

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
		"comments": &graphql.Field{
			Type: graphql.NewList(commentType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// in nested resolver, the parent object is available via p.Source
				post, ok := p.Source.(models.Post)
				if !ok {
					return nil, fmt.Errorf("Invalid Post type")
				}

				db := p.Context.Value("db").(*sqlx.DB) //we pass the Db via context

				var comments []models.Comment
				log.Printf("Fetching comments for post %d", post.Id)
				queryString := "SELECT * FROM comments WHERE post_Id = ?"
				err := db.Select(&comments, queryString, post.Id)
				if err != nil {
					return nil, err
				}
				return comments, nil

			},
		},
	},
})

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"post_id": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
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
			"posts": &graphql.Field{
				Type: graphql.NewList(postType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var posts []models.Post
					database.Select(&posts, "SELECT * FROM posts")
					return posts, nil
				},
			},
		},
	})

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"author": &graphql.ArgumentConfig{
						Type: graphql.String
					}
				},
			},
		},
	})
	cfg := graphql.SchemaConfig{
		Query: rootQuery,
	}

	return graphql.NewSchema(cfg)
}
