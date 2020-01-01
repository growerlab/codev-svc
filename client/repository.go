package client

import "fmt"

type Repository struct {
	client APISubmitter
}

func (r *Repository) Create(repo *RepoContext) (err error) {
	body := `
mutation CreateRepo() {
	createRepo(path: "%s", name: "%s") {
		name
	}
}
`
	body = fmt.Sprintf(body, repo.Path, repo.Name)
	_, err = r.client.Mutation(NewRequest(body, repo, nil))
	if err != nil {
		return
	}
	return nil
}
