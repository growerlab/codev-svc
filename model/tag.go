package model

type Tag struct {
  Name   string     `json:"name"`
  Commit *Commit    `json:"commit"`
}
