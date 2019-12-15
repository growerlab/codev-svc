package model

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type EntryType uint8

const (
	EntryTree EntryType = iota
	EntryBlob
	EntryCommit
)

type Entry struct {
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	EntryType EntryType `json:"entry_type"`

	RawEntry *plumbing.Reference
}

func InitEntry(rawEntry *plumbing.Reference) *Entry {
	entry := &Entry{RawEntry: rawEntry}
	// set Sha, Message...
	return entry
}
