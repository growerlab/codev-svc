package client

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Result struct {
	Data  json.RawMessage `json:"data"`
	Error error           `json:"error,omitempty"`
}

func (d Result) DataPath(p string) gjson.Result {
	return gjson.GetBytes(d.Data, p)
}

type APISubmitter interface {
	Query(body string, vars map[string]interface{}) (*Result, error)
	Mutation(body string, vars map[string]interface{}) (*Result, error)
}

type Repo struct {
	Name string
	Path string
}

func (r *Repo) ToVars() map[string]interface{} {
	return map[string]interface{}{
		"path": r.Path,
		"name": r.Name,
	}
}
