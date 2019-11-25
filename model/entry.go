package model

import (
  "gopkg.in/libgit2/git2go.v27"
)

type EntryType uint8

const (
  EntryTree EntryType = iota
  EntryBlob
  EntryCommit
)

type Entry struct {
  Path    string          `json:"path"`
  Name    string          `json:"name"`
  EntryType EntryType     `json:"entry_type"`

  RawEntry *git.TreeEntry
}

func InitEntry(rawEntry *git.TreeEntry) *Entry{
  entry := &Entry{RawEntry: rawEntry}
  // set Sha, Message...
  return  entry
}
