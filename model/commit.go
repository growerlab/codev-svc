package model

type Commit struct {
  SHA       string     `json:"sha"`
  Message   string     `json:"message"`
  Author    string     `json:"author"`
  committer string     `json:"committer"`
  Parent     *Commit   `json:"parent"`
  Parents []*Commit    `json:"parents"`
  Tree *Tree           `json:"tree"`
}
