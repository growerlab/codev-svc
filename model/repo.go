package model

import (
	"path"
  "gopkg.in/libgit2/git2go.v27"
)

const ReposPath = "repos/"
const DefaultBranch = "master"

type Repo struct {
  Path              string              `json:"path"`
  Name              string              `json:"name"`
  defaultBranch  *Branch                `json:"default_branch"`

  // bytes
  RepoSize          float64             `json:"repo_size"`

  Branches       []*Branch              `json:"branches"`

  Tags           []*Tag                  `json:"tags"`

  Refs           []*Ref                  `json:"refs"`

	Submodules		 []*Submodule						 `json:"submodules"`

  // internal methods
  RawRepo        *git.Repository
}

func OpenRepo(repoPath string, name string) (*Repo, error) {
  repo := &Repo{
    Path: repoPath,
    Name: name,
  }
  repoFullPath := path.Join(ReposPath, repoPath, name)
  rawRepo, err := git.OpenRepository(repoFullPath)

  if(err != nil) {
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
  repoFullPath := path.Join(ReposPath, repoPath, name )
  rawRepo, err := git.InitRepository(repoFullPath, true)
  if(err != nil) {
    return nil, err
  }
	repo.RawRepo = rawRepo
	repo.postRepoCreated()
  return repo, nil
}

func (repo *Repo)postRepoCreated() {
	// fill all fields after repo oject created

	// References
	repo.Refs = make([]*Ref, 1)
	refsIterator, err := repo.RawRepo.NewReferenceIterator()
	if(err == nil) {
		for {
			rawRef, _ := refsIterator.Next()
			if(rawRef == nil) {
				break
			}

			repo.Refs = append(repo.Refs, InitRef(rawRef.Name(), rawRef))
		}
	}


	// Branches
	repo.Branches = make([]*Branch, 1)
	branchesIterator, err := repo.RawRepo.NewBranchIterator(git.BranchLocal)
	if(err == nil) {
		for {
			rawBranch, _, _ := branchesIterator.Next()
			if(rawBranch == nil) {
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
		if(rawTag != nil) {
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

func (repo *Repo)Head()(*Ref, error) {
  rawRef, err := repo.RawRepo.Head()
  if(err != nil) {
    return nil, err
  }

  ref := &Ref{Name: rawRef.Name()}

  return ref, nil
}

func (repo *Repo)DefaultBranch() (*Branch, error) {
	rawRef, err := repo.RawRepo.Head()
	if(err != nil) {
		return nil, err
	}

	branch := &Branch{Name: rawRef.Name()}
	return branch, nil
}
