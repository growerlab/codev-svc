package model

type Commit struct {
  Parent *Commit
  Parents []*Commit
  SHA  string     `json:"sha"`
  Tree *Tree
}
