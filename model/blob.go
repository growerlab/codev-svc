package model

type Blob struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  Content string     `json:"content"`
}
