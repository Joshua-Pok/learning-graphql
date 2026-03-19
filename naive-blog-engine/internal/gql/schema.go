package gql

import (
	"github.com/Joshua-Pok/naive-blog/internal/db"
	"github.com/graphql-go/graphql"
)

var Post = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				db = db.InitDB
				db.Select("")
			},
		},
	},
})
