package utils

import "gopkg.in/src-d/go-git.v4/plumbing"

func ReferenceCompare(a, b *plumbing.Reference) bool {
	if a != nil && b != nil {
		return a.Hash() == b.Hash()
	} else if a == nil && b == nil {
		return true
	}
	return false
}
