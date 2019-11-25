package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Tree struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  Entries []*Entry   `json:"entries"`
  Trees   []*Tree    `json:"trees"`
  Blobs   []*Blob    `json:"blobs"`
  Submodules []*Submodule `json:"submodules"`

  RawTree *git.Tree
}

func InitTree(rawTree *git.Tree) *Tree{
  tree := &Tree{RawTree: rawTree}
  // set Path, Name...
  return  tree
}
