package model

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
  Path              string
  Name              string
  DefaultBranch     string
  defaultRawBranch  *Branch

  RepoSize          float64 // bytes

  Branches          []string
  RawBranches       []*Branch


  Tags              []string
  RawTags           []*Tag

  Refs              []string
  RawRefs           []*Ref
}

func (*Repo)InitRepo(path string, name string) (repo *Repo, err error) {
  repo = &Repo{
    Path: path,
    Name: name,
  }
  return
}
