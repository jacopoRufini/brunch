package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io"
	"strings"
	"time"
)

func CheckIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type Config struct {
	count int
}

func main() {
	var path string

	pathFlag := flag.String("repo", "", "the repository to use (default '.')")

	flag.Parse()

	if *pathFlag == "" {
		path = "."
	} else {
		path = *pathFlag
	}

	repo, err := git.PlainOpen(path)
	CheckIfError(err)

	refs, err := repo.Branches()
	// references should be sorted by -1
	CheckIfError(err)

	config := Config{
		count: 8,
	}

	for i := 0; i < config.count; i++ {
		r, err := refs.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			CheckIfError(err)
		}

		hash := r.Hash()
		object, err := repo.CommitObject(hash)
		CheckIfError(err)

		trim := func(s string) string {
			if len(s) < 50 {
				return s
			}

			return fmt.Sprintf("%s...\n", s[:50])
		}

		fmt.Printf("%s | %s | %s",
			fmt.Sprintf("%d days ago", int(time.Since(object.Author.When).Hours()/24)),
			strings.Replace(r.Name().String(), "refs/heads/", "", 1),
			trim(object.Message))
	}

	CheckIfError(err)
}
