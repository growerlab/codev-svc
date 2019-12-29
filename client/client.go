package client

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

const (
	defaultClientTimeout = 10 * time.Second
)

type Client struct {
	apiURL string
	client *http.Client
}

func NewClient(apiURL string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		timeout = defaultClientTimeout
	}
	_, err := url.Parse(apiURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Client{
		apiURL: apiURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) Query(body string, vars map[string]interface{}) (*Result, error) {
	bodyMap := map[string]interface{}{
		"query":     body,
		"variables": vars,
	}
	if len(vars) == 0 {
		return nil, errors.New("vars is required")
	}
	if _, found := vars["path"]; !found {
		return nil, errors.New("repo path is required in vars")
	}
	if _, found := vars["name"]; !found {
		return nil, errors.New("repo name is required in vars")
	}
	return Post(c.client, c.apiURL, bodyMap)
}

func (c *Client) Mutation(body string, vars map[string]interface{}) (*Result, error) {
	return nil, nil
}

func (c *Client) Branch(repo *Repo) *Branch {
	return &Branch{
		client: c,
		repo:   repo,
	}
}
