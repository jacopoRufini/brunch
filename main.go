package main

import (
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
)

func CheckIfError(err error) {
	if err != nil {
		panic(err)
	}
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

	refs, err := repo.References()
	CheckIfError(err)

	i := 0
	for i < 8 {
		r, err := refs.Next()
		CheckIfError(err)

		hash := r.Hash()
		object, err := repo.CommitObject(hash)
		CheckIfError(err)

		fmt.Printf("%s | %s | %s",
			object.Author.When,
			r.Name(),
			object.Message)

		i++
	}

	CheckIfError(err)
}
