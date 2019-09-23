package model

type Commit struct {
  Parent *Commit
  Parents []*Commit
  SHA  string
  Tree *Tree
}
