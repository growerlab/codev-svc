package model

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
	Path          string  `json:"path"`
	Name          string  `json:"name"`
	DefaultBranch *Branch `json:"default_branch"`

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

	rawRepo, err := git.Open(repo.fileSystem())

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

	sto, _ := repo.fileSystem()
	rawRepo, err := git.Init(sto, nil)
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
	repo.Refs = make([]*Ref, 0)
	refsIterator, err := repo.RawRepo.References()
	if err == nil {
		_ = refsIterator.ForEach(func(rawRef *plumbing.Reference) error {
			repo.Refs = append(repo.Refs, InitRef(rawRef.Name().String(), rawRef))
			return nil
		})
	}

	// Branches
	repo.Branches = make([]*Branch, 0)
	branchesIterator, err := repo.RawRepo.Branches()
	if err == nil {
		_ = branchesIterator.ForEach(func(rawBranch *plumbing.Reference) error {
			name := rawBranch.Name().String()
			repo.Branches = append(repo.Branches, InitBranch(name, rawBranch))
			return nil
		})
	}

	// Tags
	repo.Tags = make([]*Tag, 0)
	tagsInterator, err := repo.RawRepo.Tags()
	if err == nil {
		_ = tagsInterator.ForEach(func(tag *plumbing.Reference) error {
			rawTag, _ := repo.RawRepo.TagObject(tag.Hash())
			if rawTag != nil {
				repo.Tags = append(repo.Tags, InitTag(tag.Name().String(), rawTag))
			}
			return nil
		})
	}

	// Submodules
	repo.Submodules = make([]*Submodule, 1)
	tree, err := repo.RawRepo.Worktree()
	if err == nil {
		submodules, err := tree.Submodules()
		if err == nil {
			for _, sub := range submodules {
				repo.Submodules = append(repo.Submodules, InitSubmodule(sub))
			}
		}
	}
}

func (repo *Repo) Head() (*Ref, error) {
	rawRef, err := repo.RawRepo.Head()
	if err != nil {
		return nil, err
	}

	ref := &Ref{Name: rawRef.Name().String(), RawRef: rawRef}

	return ref, nil
}

func (repo *Repo) GetDefaultBranch() (*Branch, error) {
	rawRef, err := repo.RawRepo.Head()
	if err != nil {
		return nil, err
	}

	refTarget := rawRef.Target()

	if refTarget.IsBranch() || refTarget.IsTag() || refTarget.IsRemote() {

		branch := &Branch{
			Name:      rawRef.Name().String(),
			RawBranch: rawRef,
		}
		repo.DefaultBranch = branch
		return branch, nil
	}

	// raise exception right now
	return nil, errors.New("head detached")
}

func (repo *Repo) Size() int64 {
	var size int64
	err := filepath.Walk(repo.RepoPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		return 0
	}
	return size
}

func (repo *Repo) fileSystem() (storage.Storer, billy.Filesystem) {
	fs := osfs.New(repo.RepoPath)

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})

	return s, fs
}
