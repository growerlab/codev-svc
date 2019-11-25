package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Branch struct {
  Name string     `json:"name"`
  Ref  *Ref       `json:"ref"`
  Commit *Commit  `json:"commit"`

  RawBranch *git.Branch
}

func InitBranch(rawBranch *git.Branch) *Branch{
  branch := &Branch{RawBranch: rawBranch}
  // set name, Commit...
  return branch
}
