package model

import (
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	// git "gopkg.in/src-d/go-git.v4
)

type Blob struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Content string `json:"content"`

	RawBlob *object.Blob
}

func InitBlob(rawBlob *object.Blob) *Blob {
	blob := &Blob{RawBlob: rawBlob}
	// set Path, Name...
	return blob
}
