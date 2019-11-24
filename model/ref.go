package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type RefType string

const (
  RefBranch RefType = "Branch"
  RefTag RefType = "Tag"
  RefRemote RefType = "Remote"
  RefNote RefType = "Note"

  RefUnknown RefType = "Unknown"
)


type Ref struct {
  Name   string     `json:"name"`
  RefType RefType `json:"ref_type"`
  Commit *Commit    `json:"commit"`

  // internal variable
  rawRef *git.Reference
}

// func (ref*Ref) GetRefType() RefType {
//   if(ref.rawRef.IsBranch()) {
//     ref.RefType = RefBranch
//   } else if (ref.rawRef.isTag()) {
//     ref.RefType = RefTag
//   } else if (ref.rawRef.isRemote()) {
//     ref.RefType = RefRemote
//   } else if (ref.rawRef.isNote()) {
//       ref.RefType = RefNote
//   } else {
//     ref.RefType = RefUnknown
//   }
//
//   return ref.RefType
// }
