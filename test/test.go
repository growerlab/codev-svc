package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-billy.v4/osfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"

	git2 "gopkg.in/libgit2/git2go.v27"
)

// func main() {
// 	a, err := git.OpenRepository("/Users/moli/go-project/src/github.com/growerlab/codev-svc/repos/moli")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	iter, err := a.NewBranchIterator(git.BranchLocal)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	for {
// 		br, _, err := iter.Next()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(br.Name())
// 	}
//
// }

func main() {
	git2go()
	gogit()
}

const repoPath = "/Users/moli/go-project/src/github.com/growerlab/codev-svc/repos/moli"

func git2go() {
	repo, err := git2.OpenRepository(repoPath)
	if err != nil {
		panic(err)
	}

	f, _ := os.Open(os.DevNull)
	defer f.Close()

	iter, err := repo.NewBranchIterator(git2.BranchLocal)
	if err != nil {
		panic(err)
	}

	var master *git2.Branch
	err = iter.ForEach(func(branch *git2.Branch, branchType git2.BranchType) error {
		master = branch
		name, _ := branch.Name()
		fmt.Fprintln(f, name)
		return nil
	})
	if err != nil {
		panic(err)
	}

	commit, err := repo.LookupCommit(master.Target())
	if err != nil {
		panic(err)
	}
	// fmt.Fprintln(f, commit.Id().String(), commit.RawMessage())

	n := commit.ParentCount()
	for i := uint(0); i < n; i++ {
		c := commit.Parent(i)
		fmt.Fprintln(f, c.Id().String(), c.RawMessage())
	}

}

func gogit() {
	fs := osfs.New(repoPath)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
	}

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})

	repo, err := git.Open(s, fs)
	if err != nil {
		panic(err)
	}
	iter, err := repo.Branches()
	if err != nil {
		panic(err)
	}

	f, _ := os.Open(os.DevNull)
	defer f.Close()

	var master *plumbing.Reference
	err = iter.ForEach(func(reference *plumbing.Reference) error {
		master = reference
		fmt.Fprintln(f, reference.Name())
		return nil
	})
	if err != nil {
		panic(err)
	}

	c, err := repo.CommitObject(master.Hash())
	if err != nil {
		panic(err)
	}

	// fmt.Fprintln(f, c.ID().String(), c.Message)

	n := c.NumParents()
	for i := 0; i < n; i++ {
		cmt, _ := c.Parent(i)
		fmt.Fprintln(f, cmt.ID().String(), cmt.Message)
	}
}
