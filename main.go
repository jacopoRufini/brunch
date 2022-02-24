package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/manifoldco/promptui"
	"io"
	"sort"
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

type o struct {
	b *plumbing.Reference
	c *object.Commit
}

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func main() {
	peppers := []pepper{
		{Name: "Bell Pepper", HeatUnit: 0, Peppers: 0},
		{Name: "Banana Pepper", HeatUnit: 100, Peppers: 1},
		{Name: "Poblano", HeatUnit: 1000, Peppers: 2},
		{Name: "Jalapeño", HeatUnit: 3500, Peppers: 3},
		{Name: "Aleppo", HeatUnit: 10000, Peppers: 4},
		{Name: "Tabasco", HeatUnit: 30000, Peppers: 5},
		{Name: "Malagueta", HeatUnit: 50000, Peppers: 6},
		{Name: "Habanero", HeatUnit: 100000, Peppers: 7},
		{Name: "Red Savina Habanero", HeatUnit: 350000, Peppers: 8},
		{Name: "Dragon’s Breath", HeatUnit: 855000, Peppers: 9},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Pepper ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := peppers[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Spicy Level",
		Items:     peppers,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose number %d: %s\n", i+1, peppers[i].Name)
}

func maina() {
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

	lst := make([]*o, 0)
	for {
		r, err := refs.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			CheckIfError(err)
		}

		hash := r.Hash()
		c, err := repo.CommitObject(hash)
		CheckIfError(err)

		lst = append(lst, &o{r, c})
	}

	sort.Slice(lst, func(i, j int) bool {
		return lst[i].c.Author.When.Unix() > lst[j].c.Author.When.Unix()
	})

	branches := make([]string, 0)
	for i := 0; i < config.count; i++ {
		obj := lst[i]
		branches = append(branches, fmt.Sprintf("%s | %s\n",
			fmt.Sprintf("%d days ago", int(time.Since(obj.c.Author.When).Hours()/24)),
			strings.Replace(obj.b.Name().String(), "refs/heads/", "", 1)))
	}

	prompt := promptui.Select{
		Label: "Branch",
		Items: branches,
	}

	_, result, err := prompt.Run()
	CheckIfError(err)

	fmt.Printf("You choose %q\n", result)
}
