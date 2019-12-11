package model

import (
	"gopkg.in/libgit2/git2go.v27"
)

type Branch struct {
	Name   string  `json:"name"`
	Ref    *Ref    `json:"ref"`
	Commit *Commit `json:"commit"`

	RawBranch *git.Branch
}

func InitBranch(name string, rawBranch *git.Branch) *Branch {
	branch := &Branch{Name: name, RawBranch: rawBranch}
	// set name, Commit...
	return branch
}
