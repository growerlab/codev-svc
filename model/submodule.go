package model

import "gopkg.in/src-d/go-git.v4"

type Submodule struct {
	RawSubmodule *git.Submodule
}

func InitSubmodule(rawSubmodule *git.Submodule) *Submodule {
	submodule := &Submodule{RawSubmodule: rawSubmodule}
	return submodule
}
