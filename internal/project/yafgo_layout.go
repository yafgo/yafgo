package project

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/yafgo/yafgo/internal/pkg/file"
)

type TplYafgoLayout struct {
	repo string
}

const (
	TOY_LAYOUT_REPO_GITHUB = "https://github.com/yafgo/yafgo-layout.git"
	TOY_LAYOUT_REPO_GITEE  = "https://gitee.com/yafgo/yafgo-layout.git"
)

func (rp *TplYafgoLayout) Name() string {
	return "YafgoLayout"
}

// MakeProject
//
//	name: like "github.com/you/project" or "my_project"
func (rp *TplYafgoLayout) MakeProject(name string) (err error) {

	// select repo
	rp.repo, err = rp.selectRepo()
	if err != nil {
		return
	}

	// eg: "project" "my_project"
	projectName := path.Base(name)
	// workDir
	cwd, _ := os.Getwd()
	projectDir := path.Join(cwd, projectName)

	defer func() {
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(color.Success.Sprint("创建成功: "), projectName)
	}()

	if file.Exists(projectDir) {
		fmt.Println(color.Warn.Sprint("项目已存在: "), projectName)
		err = errors.New("project already exists")
		return
	}

	// clone project
	{
		fmt.Println(color.Gray.Sprint("创建项目: "), projectName)
		cmd := exec.Command("git", "clone", rp.repo, projectName)
		color.Grayln(cmd.String())
		err = cmd.Run()
		if err != nil {
			color.Redf("创建失败, 请检查 [%v] 访问状况并重试!\n", rp.repo)
			return
		}
	}

	// delete .git
	_ = os.RemoveAll(path.Join(projectDir, ".git"))
	_ = os.RemoveAll(path.Join(projectDir, "LICENSE"))

	// replace moduleName
	rp.renameModule(name, projectDir)

	// chdir
	if _err := os.Chdir(projectDir); _err != nil {
		err = _err
		return
	}

	// gofmt
	{
		cmd := exec.Command("gofmt", "-w", ".")
		color.Grayln(cmd.String())
		err = cmd.Run()
		if err != nil {
			color.Errorln("gofmt fail")
			return
		}
	}

	// git init
	{
		cmd := exec.Command("git", "init")
		color.Grayln(cmd.String())
		err = cmd.Run()
		if err != nil {
			color.Errorln("git init fail")
			return
		}
	}

	return
}

func (rp *TplYafgoLayout) selectRepo() (repo string, err error) {
	repos := []struct {
		Name string
		Desc string
		Repo string
	}{
		{Name: "[Github]", Desc: "github.com", Repo: TOY_LAYOUT_REPO_GITHUB},
		{Name: "[Gitee] ", Desc: "gitee.com", Repo: TOY_LAYOUT_REPO_GITEE},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }} ({{ .Desc | red }}) - {{ .Repo | magenta }}",
		Inactive: "  {{ .Name | cyan }} ({{ .Desc | red }})",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
	}

	prompt := promptui.Select{
		Label:     "Select Repo",
		Items:     repos,
		Templates: templates,
	}

	i, _, _err := prompt.Run()

	if _err != nil {
		err = _err
		return
	}

	repo = repos[i].Repo
	return
}

func (rp *TplYafgoLayout) renameModule(moduleName, dir string) (err error) {

	err = file.WalkFiles(dir, func(elem string) error {
		if strings.HasSuffix(elem, ".go") || path.Base(elem) == "go.mod" {
			// *.go, go.mod
			err := file.ReplaceString(elem, "yafgo/yafgo-layout", moduleName)
			return err
		}

		return nil
	})

	return
}
