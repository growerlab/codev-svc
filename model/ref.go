package model

type RefType uint8

const (
  RefBranch RefType = iota
  RefTag
  RefRemote
  RefNote

  RefUnknown
)


type Ref struct {
  Name   string     `json:"name"`
  RefType string `json:"ref_type"`
  Commit *Commit    `json:"commit"`

  // internal variable
  rawRef
}

func (ref*Ref) RefType() RefType {
  if(ref.rawRef.IsBranch()) {
    ref.RefType = RefType.RefBranch
  } else if (ref.rawRef.isTag()) {
    ref.RefType = RefType.RefTag
  } else if (ref.rawRef.isRemote()) {
    ref.RefType = RefType.RefRemote
  } else if (ref.rawRef.isNote()) {
      ref.RefType = RefType.RefNote
  } else {
    ref.RefType = RefType.RefUnknown
  }

  return ref.RefType
}
