package client

import "github.com/pkg/errors"

type Branch struct {
	client APISubmitter
	repo   *Repo
}

func (b *Branch) Default() (string, error) {
	body := `
{
	repo {
		default_branch {
			name
		}
	}
}`
	result, err := b.client.Query(body, b.repo.ToVars())
	if err != nil {
		return "", errors.WithStack(err)
	}
	branchName := result.DataPath("repo.default_branch.name").String()
	return branchName, nil
}
