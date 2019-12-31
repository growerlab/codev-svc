package client

type Repository struct {
	client APISubmitter
}

func (r *Repository) Create(repo *RepoContext) (err error) {
	body := `
mutation CreateRepo($path: String!, $name: String!) {
	createRepo(path: $path, name: $name) {
		name
	}
}
`
	_, err = r.client.Mutation(NewRequest(body, repo, nil))
	if err != nil {
		return
	}
	return nil
}
