package model

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
  Path              string              `json:"path"`
  Name              string              `json:"name"`
  DefaultBranch     string              `json:"default_branch"`
  defaultRawBranch  *Branch

  // bytes
  RepoSize          float64             `json:"repo_size"`

  Branches          []string            `json:"branches"`
  RawBranches       []*Branch


  Tags              []string            `json:"tags"`
  RawTags           []*Tag

  Refs              []string            `json:"refs"`
  RawRefs           []*Ref
}

func (*Repo)InitRepo(path string, name string) (repo *Repo, err error) {
  repo = &Repo{
    Path: path,
    Name: name,
  }
  return
}
