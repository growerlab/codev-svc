package model

type Branch struct {
  Name string     `json:"name"`
  Ref  *Ref
  Commit *Commit
}
