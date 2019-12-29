package client

import "testing"

func TestBranch_Default(t *testing.T) {
	client, _ := NewClient("http://localhost:9000/graphql", 0)
	defaultBranch, err := client.Branch(&RepoContext{
		Path: "/",
		Name: "moli",
	}).Default()

	if err != nil {
		t.Fatalf("%+v", err)
	}
	if defaultBranch != "master" {
		t.Fatal(defaultBranch)
	}
}
