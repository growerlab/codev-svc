package model

type Branch struct {
  Name string
  Ref  *Ref
  Commit *Commit
}
