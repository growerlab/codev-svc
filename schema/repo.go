package schema

import (
	"github.com/growerlab/letsgit-svc/model"
	"github.com/graphql-go/graphql"
)

var RepoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Repo",
	Description: "Repo Model",
	Fields: graphql.Fields{
		"Path": &graphql.Field{
			Type: graphql.String,
		},
		"Name": &graphql.Field{
			Type: graphql.String,
		},
		"RepoSize": &graphql.Field{
			Type: graphql.Float,
		},
		"DefaultBranch": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryRepo = graphql.Field{
	Name: "repo",
	Description: "Query Repo",
	Type: graphql.NewNonNull(RepoType),
	Args: graphql.FieldConfigArgument{
		"Path": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"Name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		path, _ := p.Args["Path"].(string)
		name, _ := p.Args["Name"].(string)

		return (&model.Repo{}).InitRepo(path, name)
	},
}
