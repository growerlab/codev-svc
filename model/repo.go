package model

import (
	"path"
  "gopkg.in/libgit2/git2go.v27"
)

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
  Path              string              `json:"path"`
  Name              string              `json:"name"`
  defaultBranch  *Branch                `json:"default_branch"`

  // bytes
  RepoSize          float64             `json:"repo_size"`

  Branches       []*Branch              `json:"branches"`

  Tags           []*Tag                  `json:"tags"`

  Refs           []*Ref                  `json:"refs"`

  // internal methods
  RawRepo        *git.Repository
}

func OpenRepo(repoPath string, name string) (*Repo, error) {
  repo := &Repo{
    Path: repoPath,
    Name: name,
  }
  repoFullPath := path.Join(ReposPath, repoPath, name)
  rawRepo, err := git.OpenRepository(repoFullPath)

  if(err != nil) {
    return nil, err
  }
	repo.RawRepo = rawRepo

  return repo, nil
}

func InitRepo(repoPath string, name string) (*Repo, error) {
  repo := &Repo{
    Path: repoPath,
    Name: name,
  }
  repoFullPath := path.Join(ReposPath, repoPath, name )
  rawRepo, err := git.InitRepository(repoFullPath, true)
  if(err != nil) {
    return nil, err
  }
	repo.RawRepo = rawRepo
  return repo, nil
}

func (repo *Repo)Head()(*Ref, error) {
  rawRef, err := repo.RawRepo.Head()
  if(err != nil) {
    return nil, err
  }

  ref := &Ref{Name: rawRef.Name()}

  return ref, nil
}

func (repo *Repo)DefaultBranch() (*Branch, error) {
	rawRef, err := repo.RawRepo.Head()
	if(err != nil) {
		return nil, err
	}

	branch := &Branch{Name: rawRef.Name()}
	return branch, nil
}
