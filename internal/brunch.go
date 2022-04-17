package internal

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func (i item) Timestamp() int64 {
	return i.commit.Author.When.Unix()
}

type item struct {
	branch *plumbing.Reference
	commit *object.Commit
}

// findCurrentRepository is a utility to read the repository on the current directory
// if no git repository is found, this function recursively tries to find a repository from the parent folder, until the home directory is reached
func findCurrentRepository() (*git.Repository, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get the user directory while trying to find the current git repository: %v", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("could not current directory: %v", err)
	}

	for {
		repo, err := git.PlainOpen(dir)
		if err == nil {
			return repo, nil
		}

		if dir == home {
			return nil, fmt.Errorf("could not find any git repository")
		}

		dir = filepath.Dir(dir)
	}
}

func Brunch(count int) ([]DisplayObject, error) {
	repo, err := findCurrentRepository()
	if err != nil {
		return nil, err
	}

	branches, err := repo.Branches()
	if err != nil {
		return nil, err
	}

	items := make([]item, 0)
	err = branches.ForEach(func(branch *plumbing.Reference) error {
		commit, err := repo.CommitObject(branch.Hash())
		if err != nil {
			return err
		}

		items = append(items, item{branch, commit})

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Timestamp() > items[j].Timestamp()
	})

	displayObjects := make([]DisplayObject, 0)

	until := len(items)
	if until > count {
		until = count
	}

	for _, i := range items[:until] {
		branchName := strings.Replace(i.branch.Name().String(), "refs/heads/", "", 1)
		timestamp := fmt.Sprintf("%d days ago", int(time.Since(i.commit.Author.When).Hours()/24))
		commitMessage := i.commit.Message

		displayObjects = append(displayObjects, DisplayObject{
			When:          timestamp,
			BranchName:    branchName,
			CommitMessage: commitMessage,
		})
	}

	return displayObjects, nil
}
