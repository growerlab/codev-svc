package model

type Ref struct {
  name   string     `json:"path"`
  Commit *Commit
}
