package schema

import "github.com/graphql-go/graphql"

var branchType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Repo",
	Description: "Repo Model",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"commits": &graphql.Field{
			Type: graphql.String,
		},
	},
})
