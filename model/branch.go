package model

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Branch struct {
	Name   string  `json:"name"`
	Ref    *Ref    `json:"ref"`
	Commit *Commit `json:"commit"`

	RawBranch *plumbing.Reference
}

func InitBranch(name string, rawBranch *plumbing.Reference) *Branch {
	branch := &Branch{Name: name, RawBranch: rawBranch}
	// set name, Commit...
	return branch
}
