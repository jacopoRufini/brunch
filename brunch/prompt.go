package brunch

import "fmt"

func (d DisplayObject) String() string {
	return d.BranchName + "\n"
}

type DisplayObject struct {
	DisplayName string
	BranchName  string
}

func Prompt(objects []DisplayObject) {
	fmt.Println(objects)
	//	templates := &promptui.SelectTemplates{
	//		Label:    "{{ . }}?",
	//		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .HeatUnit | red }})",
	//		Inactive: "  {{ .Name | cyan }} ({{ .HeatUnit | red }})",
	//		Selected: "\U0001F336 {{ .Name | red | cyan }}",
	//		Details: `
	//--------- Pepper ----------
	//{{ "Name:" | faint }}	{{ .Name }}
	//{{ "Heat Unit:" | faint }}	{{ .HeatUnit }}
	//{{ "Peppers:" | faint }}	{{ .Peppers }}`,
	//	}
	//
	//	searcher := func(input string, index int) bool {
	//		object := objects[index]
	//		name := strings.Replace(strings.ToLower(object.BranchName), " ", "", -1)
	//		input = strings.Replace(strings.ToLower(input), " ", "", -1)
	//
	//		return strings.Contains(name, input)
	//	}
	//
	//	prompt := promptui.Select{
	//		Label:     "Branches",
	//		Items:     objects,
	//		Templates: templates,
	//		Size:      len(objects),
	//		Searcher:  searcher,
	//	}
	//
	//	i, _, err := prompt.Run()
	//
	//	if err != nil {
	//		fmt.Printf("Prompt failed %v\n", err)
	//		return
	//	}
	//
	// fmt.Printf("You choose number %d: %s\n", i+1, objects[i].BranchName)
}
