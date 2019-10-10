package model

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
}
