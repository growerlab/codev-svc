package model

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/growerlab/codev-svc/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var ReposDir = "repos/"

const hooksDir = "template/hooks"
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

	// Submodules []*Submodule `json:"submodules"`

	// internal methods
	RawRepo  *git.Repository
	RepoPath string
}

func OpenRepo(repoPath string, name string) (*Repo, error) {
	repo := &Repo{
		Path: repoPath,
		Name: name,
	}
	log.Info().Str("path", repo.RepoPath).Msg("open repo")
	repo.RepoPath = path.Join(ReposDir, repoPath, name)

	rawRepo, err := git.PlainOpen(repo.RepoPath)

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
	repo.RepoPath = path.Join(ReposDir, repoPath, name)

	rawRepo, err := git.PlainInit(repo.RepoPath, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	repo.RawRepo = rawRepo
	err = repo.initHooks()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	repo.postRepoCreated()
	return repo, nil
}

// TODO 应为push操作增加软链hook
func (repo *Repo) initHooks() error {
	hookInfos, err := ioutil.ReadDir(hooksDir)
	if err != nil {
		return errors.WithStack(err)
	}

	newHookDir := filepath.Join(repo.RepoPath, "hooks")
	if _, err := os.Stat(newHookDir); os.IsNotExist(err) {
		err = os.Mkdir(newHookDir, 0755)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	for _, hook := range hookInfos {
		hookPath := filepath.Join(hooksDir, hook.Name())
		newHookPath := filepath.Join(newHookDir, hook.Name())
		err = utils.CopyFile(hookPath, newHookPath, 0755)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
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
	repo.Branches, err = repo.branches()
	if err != nil {
		return
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
	// repo.Submodules = make([]*Submodule, 0)
	// tree, err := repo.RawRepo.Worktree()
	// if err == nil {
	// 	submodules, err := tree.Submodules()
	// 	if err == nil {
	// 		for _, sub := range submodules {
	// 			repo.Submodules = append(repo.Submodules, InitSubmodule(sub))
	// 		}
	// 	}
	// }
}

func (repo *Repo) Head() (*Ref, error) {
	rawRef, err := repo.RawRepo.Head()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	ref := &Ref{Name: rawRef.Name().String(), RawRef: rawRef}

	return ref, nil
}

func (repo *Repo) FileEntries(path string, hash plumbing.Hash) ([]object.TreeEntry, error) {
	tree, err := repo.RawRepo.TreeObject(hash)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pathTree, err := tree.Tree(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return pathTree.Entries, nil
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

func (repo *Repo) branches() ([]*Branch, error) {
	if len(repo.Branches) > 0 {
		return repo.Branches, nil
	}

	iter, err := repo.RawRepo.Branches()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	branes := make([]*Branch, 0)
	err = iter.ForEach(func(ref *plumbing.Reference) error {
		branes = append(branes, InitBranch(ref.Name().String(), ref))
		return nil
	})
	repo.Branches = branes
	return repo.Branches, errors.WithStack(err)
}

func (repo *Repo) CreateBranch(name string) error {
	err := repo.RawRepo.CreateBranch(&config.Branch{
		Name: name,
	})
	return errors.WithStack(err)
}

func (repo *Repo) DeleteBranch(name string) error {
	err := repo.RawRepo.DeleteBranch(name)
	return errors.WithStack(err)
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
