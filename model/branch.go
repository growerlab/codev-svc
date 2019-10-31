package model

type Branch struct {
  Name string     `json:"name"`
  Ref  *Ref       `json:"ref"`
  Commit *Commit  `json:"commit"`
}
