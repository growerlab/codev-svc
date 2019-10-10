package model

type Tag struct {
  name   string     `json:"path"`
  Commit *Commit 
}
