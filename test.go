package main

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4/plumbing/object"

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
	// gogit()
}

const repoPath = "/Users/moli/go-project/src/github.com/growerlab/codev-svc/repos/moli"

func git2go() {
	repo, err := git2.OpenRepository(repoPath)
	if err != nil {
		panic(err)
	}

	iter, err := repo.NewBranchIterator(git2.BranchLocal)
	if err != nil {
		panic(err)
	}

	err = iter.ForEach(func(branch *git2.Branch, branchType git2.BranchType) error {
		fmt.Println(branch.Name())
		return nil
	})
	if err != nil {
		panic(err)
	}

	wk, err := repo.Walk()
	if err != nil {
		panic(err)
	}
	// wk.Iterate(func(commit *git2.Commit) bool {
	// 	fmt.Println(commit.Author().Email)
	// 	return true
	// })
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

	err = iter.ForEach(func(reference *plumbing.Reference) error {
		fmt.Println(reference.Name())
		return nil
	})
	if err != nil {
		panic(err)
	}

	commits, err := repo.CommitObjects()
	if err != nil {
		panic(err)
	}
	commits.ForEach(func(commit *object.Commit) error {
		fmt.Println(commit.Author.Email)
		return nil
	})
}
