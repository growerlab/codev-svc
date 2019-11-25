package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Blob struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  Content string     `json:"content"`

  RawBlob *git.Blob
}

func InitBlob(rawBlob *git.Blob) *Blob{
  blob := &Blob{RawBlob: rawBlob}
  // set Path, Name...
  return blob
}
