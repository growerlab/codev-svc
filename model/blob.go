package model

type Blob struct {
  Path    string     `json:"path"`
  Name    string     `json:"name"`
  Raw     []byte
  Content string     `json:"content"`
}
