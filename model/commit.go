package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Commit struct {
  Sha       string     `json:"sha"`
  Message   string     `json:"message"`
  Author    string     `json:"author"`
  committer string     `json:"committer"`
  Parent     *Commit   `json:"parent"`
  Parents []*Commit    `json:"parents"`
  Tree *Tree           `json:"tree"`

  RawCommit *git.Commit
}

func InitCommit(rawCommit *git.Commit) *Commit{
  commit := &Commit{RawCommit: rawCommit}
  // set Sha, Message...
  return  commit
}
