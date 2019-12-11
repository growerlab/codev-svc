package model

import (
	"errors"
	"gopkg.in/libgit2/git2go.v27"
)

type RefType string

const (
	RefBranch RefType = "Branch"
	RefTag    RefType = "Tag"
	RefRemote RefType = "Remote"
	RefNote   RefType = "Note"

	RefUnknown RefType = "Unknown"
)

type Ref struct {
	Name    string  `json:"name"`
	RefType RefType `json:"ref_type"`
	Commit  *Commit `json:"commit"`

	RawRef *git.Reference
	Repo   *Repo
}

func InitRef(name string, rawRef *git.Reference) *Ref {
	return &Ref{Name: name, RawRef: rawRef}
}

func (ref *Ref) RetriveRefType() RefType {
	if ref.RawRef.IsBranch() {
		ref.RefType = RefBranch
	} else if ref.RawRef.IsTag() {
		ref.RefType = RefTag
	} else if ref.RawRef.IsRemote() {
		ref.RefType = RefRemote
	} else if ref.RawRef.IsNote() {
		ref.RefType = RefNote
	} else {
		ref.RefType = RefUnknown
	}

	return ref.RefType
}

func (ref *Ref) TargetCommit() (*Commit, error) {
	refType := ref.RawRef.Type()
	if git.ReferenceSymbolic == refType {
		refName := ref.RawRef.SymbolicTarget()
		refs := ref.Repo.RawRepo.References
		refResolved, err := refs.Lookup(refName)
		if err != nil {
			return nil, err
		}
		refWrapped := &Ref{Name: refName, RawRef: refResolved}
		return refWrapped.TargetCommit()
	} else if git.ReferenceOid == refType {
		oid := ref.RawRef.Target()
		rawCommit, err := ref.Repo.RawRepo.LookupCommit(oid)
		if err != nil {
			return nil, err
		}
		return InitCommit(rawCommit), nil
	} else {
		return nil, errors.New("not found")
	}
}
