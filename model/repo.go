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

func OpenRepo(path string, name string) (*Repo, error) {
  repo := &Repo{
    Path: path,
    Name: name,
  }
  repoPath := path.Join(ReposPath, path, name )
  repo.RawRepo, err := git.OpenRepository(repoPath)

  if(err == nil) {
    return nil, err
  }

  return repo, nil
}

func InitRepo(path string, name string) (*Repo, error) {
  repo = &Repo{
    Path: path,
    Name: name,
  }
  repoPath := path.Join(ReposPath, path, name )
  repo.RawRepo, err := git.InitRepository(repoPath, true)
  if(err == nil) {
    return nil, err
  }
  return repo, nil
}

func (repo *Repo)Head()(*Ref, err) {
  rawRef, err := repo.Head()
  if(err == nil) {
    return nil, err
  }


  ref := &Ref{name: rawRef.Name()}

  return ref, nil
}

func (repo *Repo)DefaultBranch(*Branch, err) {
    ref, err := repo.Head()
    if(ref == nil) {
			return nil, nil
		}
		return nil, nil
}
