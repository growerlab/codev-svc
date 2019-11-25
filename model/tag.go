package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Tag struct {
  Name   string     `json:"name"`
  Commit *Commit    `json:"commit"`

  RawTag *git.Tag
}

func InitTag(name string, rawTag *git.Tag) *Tag{
  return &Tag{Name: name, RawTag: rawTag}
}
