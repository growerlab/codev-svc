package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type Tag struct {
  Name   string     `json:"name"`
  Commit *Commit    `json:"commit"`

  RawTag *git.Tag
}

func InitTag(rawTag *git.Tag) *Tag{
  tag := &Tag{RawTag: rawTag}
  // set Name, Commit...
  return  tag
}
