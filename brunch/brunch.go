package brunch

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
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

func Brunch(repositoryName string, count int) ([]DisplayObject, error) {
	repo, err := git.PlainOpen(repositoryName)
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
