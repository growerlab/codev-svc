package model

type Tree struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  Entries []*Entry   `json:"entries"`
  Trees   []*Tree    `json:"trees"`
  Blobs   []*Blob    `json:"blobs"`
  Submodules []*Submodule `json:"submodules"`
}
