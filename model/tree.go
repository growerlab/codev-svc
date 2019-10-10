package model

type Tree struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  entries []*Entry
}
