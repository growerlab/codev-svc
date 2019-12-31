package client

import "testing"

func TestBranchInfo(t *testing.T) {
	client, repo := defaultClient()
	defaultBranch, branches, err := client.Branch(repo).Info()
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if defaultBranch != "master" {
		t.Fail()
	}

	if len(branches) != len([]string{"master"}) {
		t.Fatal(branches)
	}
}

func defaultClient() (*Client, *RepoContext) {
	client, _ := NewClient("http://localhost:9000/graphql", 0)
	repo := &RepoContext{
		Path: "/",
		Name: "moli",
	}
	return client, repo
}
