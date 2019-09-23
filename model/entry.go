package model

type EntryType uint8

const (
  EntryTree EntryType = iota
  EntryBlob
  EntryCommit
)

type Entry struct {
  Path    string
  Name    string
  EntryType EntryType
}
