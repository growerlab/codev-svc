package schema

import (
	"github.com/growerlab/codev-svc/model"
	"github.com/graphql-go/graphql"
)

var RepoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Repo",
	Description: "Repo Model",
	Fields: graphql.Fields{
		"path": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"pepo_size": &graphql.Field{
			Type: graphql.Float,
		},
		"default_branch": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryRepo = graphql.Field{
	Name: "repo",
	Description: "Query Repo",
	Type: graphql.NewNonNull(RepoType),
	Args: graphql.FieldConfigArgument{
		"path": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
		path, _ := p.Args["path"].(string)
		name, _ := p.Args["name"].(string)

		repo, err := model.InitRepo(path, name)
		if(err != nil) {
			return nil, err
		}
		return repo, nil
	},
}
