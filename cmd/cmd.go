package cmd

import (
	"errors"
	"strings"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/yafgo/yafgo/consts"
	"github.com/yafgo/yafgo/internal/project"
)

// Tpl Project Layout Template
type Tpl interface {
	Name() string
	MakeProject(name string) error
}

// Execute make project
func Execute() {

	color.Successf("[Yafgo] v%s\n\n", consts.Version)

	// eg: "github.com/you/project"  "my_project"
	moduleName, _ := readProjectName()
	if moduleName == "" {
		// fmt.Println("cancel")
		return
	}

	tpl, err := selectTemplate()
	if err != nil || tpl == nil {
		color.Grayf("select template: %v\n", err)
		return
	}

	err = tpl.MakeProject(moduleName)
	if err != nil {
		color.Errorf("make project: %v\n", err)
		return
	}
}

func readProjectName() (name string, err error) {
	validate := func(input string) error {
		if len(strings.TrimSpace(input)) < 3 {
			return errors.New("the Project name must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Project Name",
		Validate: validate,
		Default:  "my_project",
	}

	name, err = prompt.Run()

	if err != nil {
		return
	}

	name = strings.TrimSpace(name)
	name = strings.ToLower(name)

	return
}

func selectTemplate() (tpl Tpl, err error) {
	projects := []struct {
		Name string
		Desc string
		Flag string
	}{
		{Name: "[Yafgo]   ", Desc: "Yafgo Layout", Flag: "yafgo"},
		{Name: "[Goravel] ", Desc: "Goravel", Flag: "goravel"},
	}

	promptTpls := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .Desc | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Desc | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select Template",
		Items:     projects,
		Templates: promptTpls,
	}

	i, _, _err := prompt.Run()

	if _err != nil {
		err = _err
		return
	}

	switch projects[i].Flag {
	case "yafgo":
		tpl = new(project.TplYafgoLayout)
	case "goravel":
		tpl = new(project.TplGoravel)
	}

	return
}
