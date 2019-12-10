package model

import (
	"errors"
	"gopkg.in/libgit2/git2go.v27"
	"os"
	"path"
	"path/filepath"
)

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
	Path          string  `json:"path"`
	Name          string  `json:"name"`
	defaultBranch *Branch `json:"default_branch"`

	// bytes
	RepoSize float64 `json:"repo_size"`

	Branches []*Branch `json:"branches"`

	Tags []*Tag `json:"tags"`

	Refs []*Ref `json:"refs"`

	Submodules []*Submodule `json:"submodules"`

	// internal methods
	RawRepo  *git.Repository
	RepoPath string
}

func OpenRepo(repoPath string, name string) (*Repo, error) {
	repo := &Repo{
		Path: repoPath,
		Name: name,
	}
	repo.RepoPath = path.Join(ReposPath, repoPath, name)
	rawRepo, err := git.OpenRepository(repo.RepoPath)

	if err != nil {
		return nil, err
	}
	repo.RawRepo = rawRepo
	repo.postRepoCreated()

	return repo, nil
}

func InitRepo(repoPath string, name string) (*Repo, error) {
	repo := &Repo{
		Path: repoPath,
		Name: name,
	}
	repo.RepoPath = path.Join(ReposPath, repoPath, name)
	rawRepo, err := git.InitRepository(repo.RepoPath, true)
	if err != nil {
		return nil, err
	}
	repo.RawRepo = rawRepo
	repo.postRepoCreated()
	return repo, nil
}

func (repo *Repo) postRepoCreated() {
	// fill all fields after repo oject created

	// References
	repo.Refs = make([]*Ref, 1)
	refsIterator, err := repo.RawRepo.NewReferenceIterator()
	if err == nil {
		for {
			rawRef, _ := refsIterator.Next()
			if rawRef == nil {
				break
			}

			repo.Refs = append(repo.Refs, InitRef(rawRef.Name(), rawRef))
		}
	}

	// Branches
	repo.Branches = make([]*Branch, 1)
	branchesIterator, err := repo.RawRepo.NewBranchIterator(git.BranchLocal)
	if err == nil {
		for {
			rawBranch, _, _ := branchesIterator.Next()
			if rawBranch == nil {
				break
			}

			name, _ := rawBranch.Name()
			repo.Branches = append(repo.Branches, InitBranch(name, rawBranch))
		}
	}

	// Tags
	repo.Tags = make([]*Tag, 1)
	repo.RawRepo.Tags.Foreach(func(name string, oid *git.Oid) error {
		rawTag, _ := repo.RawRepo.LookupTag(oid)
		if rawTag != nil {
			repo.Tags = append(repo.Tags, InitTag(name, rawTag))
		}
		return nil
	})

	// Submodules
	repo.Submodules = make([]*Submodule, 1)
	repo.RawRepo.Submodules.Foreach(func(rawSubmodule *git.Submodule, name string) int {
		repo.Submodules = append(repo.Submodules, InitSubmodule(rawSubmodule))
		return 1
	})
}

func (repo *Repo) Head() (*Ref, error) {
	rawRef, err := repo.RawRepo.Head()
	if err != nil {
		return nil, err
	}

	ref := &Ref{Name: rawRef.Name(), RawRef: rawRef}

	return ref, nil
}

func (repo *Repo) DefaultBranch() (*Branch, error) {
	rawRef, err := repo.RawRepo.Head()
	if err != nil {
		return nil, err
	}

	if rawRef.IsBranch() || rawRef.IsTag() || rawRef.IsRemote() {
		branch := &Branch{Name: rawRef.Name(), RawBranch: rawRef.Branch()}
		return branch, nil
	}

	// raise exception right now
	return nil, errors.New("head detached")
}

func (repo *Repo) Size() int64 {
	var size int64
	filepath.Walk(repo.RepoPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}
