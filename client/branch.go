package client

import "github.com/pkg/errors"

type Branch struct {
	client APISubmitter
	repo   *RepoContext
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
	result, err := b.client.Query(NewRequest(body, b.repo, nil))
	if err != nil {
		return "", errors.WithStack(err)
	}
	branchName := result.DataPath("repo.default_branch.name").String()
	return branchName, nil
}
