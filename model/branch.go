package model

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type Branch struct {
	Default bool    `json:"default"`
	Name    string  `json:"name"`
	Ref     *Ref    `json:"ref"`
	Commit  *Commit `json:"commit"`

	RawBranch *plumbing.Reference
}

func InitBranch(rawBranch *plumbing.Reference) *Branch {
	name := rawBranch.Name().Short()
	branch := &Branch{Name: name, RawBranch: rawBranch}
	// set name, Commit...
	return branch
}
