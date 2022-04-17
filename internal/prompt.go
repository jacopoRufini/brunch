package internal

import (
	"github.com/manifoldco/promptui"
)

type DisplayObject struct {
	When          string
	BranchName    string
	CommitMessage string
}

func Prompt(objects []DisplayObject) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "⇨ {{ .BranchName | yellow }}",
		Inactive: "  {{ .BranchName | cyan }}",
		Selected: "✘ {{ .BranchName | yellow }}",
		Details: `
--------- INFO ----------
{{ "When:" | faint }}	{{ .When }},
{{ "Message:" | faint }}	{{ .CommitMessage }}`,
	}

	prompt := promptui.Select{
		Label:     "Select branch",
		Items:     objects,
		Size:      8,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return objects[i].BranchName, nil
}
