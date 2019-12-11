package model

import (
	"gopkg.in/libgit2/git2go.v27"
)

type Submodule struct {
	RawSubmodule *git.Submodule
}

func InitSubmodule(rawSubmodule *git.Submodule) *Submodule {
	submodule := &Submodule{RawSubmodule: rawSubmodule}
	return submodule
}
