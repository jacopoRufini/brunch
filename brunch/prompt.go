package brunch

import (
	"github.com/manifoldco/promptui"
	"strings"
)

type DisplayObject struct {
	When          string
	BranchName    string
	CommitMessage string
}

func Prompt(objects []DisplayObject) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .BranchName | cyan }}",
		Inactive: "  {{ .BranchName | cyan }}",
		Selected: "\U0001F336 {{ .BranchName | red | cyan }}",
		Details: `
--------- INFO ----------
{{ "When:" | faint }}	{{ .When }},
{{ "Message:" | faint }}	{{ .CommitMessage }}`,
	}

	searcher := func(input string, index int) bool {
		obj := objects[index]
		name := strings.Replace(strings.ToLower(obj.BranchName), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select branch",
		Items:     objects,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return objects[i].BranchName, nil
}
