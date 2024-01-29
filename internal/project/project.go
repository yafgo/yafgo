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

type TplBase struct {
	repo       string
	moduleName string // go项目名称(go module的模块名称, 如: "github.com/you/project")
	dirName    string // 项目目录名(本地目录名, 如: "project")
	projectDir string // 项目路径(本地路径, 如: "/path/to/project")

	REPO_GITHUB string
	REPO_GITEE  string
}

// initProjectDir 初始化项目名称、项目目录、项目路径
//
//	name: like "github.com/you/project" or "my_project"
func (p *TplBase) initProjectDir(name string) (err error) {
	// go项目名称
	p.moduleName = name
	// 目录名称, 如: "project" "my_project"
	p.dirName = path.Base(name)

	// workDir
	cwd, _ := os.Getwd()
	p.projectDir = path.Join(cwd, p.dirName)

	if file.Exists(p.projectDir) {
		fmt.Println(color.Warn.Sprint("文件已存在: "), p.dirName)
		err = errors.New("project already exists")
		return
	}
	return
}

// selectRepo 选择仓库源
func (p *TplBase) selectRepo() (err error) {
	repos := []struct {
		Name string
		Desc string
		Repo string
	}{
		{Name: "[Github]", Desc: "github.com", Repo: p.REPO_GITHUB},
		{Name: "[Gitee] ", Desc: "gitee.com", Repo: p.REPO_GITEE},
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

	p.repo = repos[i].Repo
	return
}

// cloneProject 克隆项目
func (p *TplBase) cloneProject() (err error) {
	fmt.Println(color.Gray.Sprint("创建项目: "), p.dirName)

	cmdArgs := []string{"clone"}
	// 分支或tag
	var branchName string
	branchName, err = p.readBranchOrTagName()
	if err != nil {
		return
	}
	if branchName != "" {
		cmdArgs = append(cmdArgs, "-b", branchName)
	}
	cmdArgs = append(cmdArgs, "--depth", "1", p.repo, p.dirName)

	// 克隆项目
	cmd := exec.Command("git", cmdArgs...)
	color.Grayln(cmd.String())
	err = cmd.Run()
	if err != nil {
		color.Redf("创建失败, 请检查 [%v] 访问状况并重试!\n", p.repo)
		return
	}

	return
}

func (p *TplBase) runGoFmt() (err error) {
	cmd := exec.Command("gofmt", "-w", ".")
	color.Grayln(cmd.String())
	err = cmd.Run()
	if err != nil {
		color.Errorln("gofmt fail")
		return
	}
	return
}

func (p *TplBase) runGitInit() (err error) {
	cmd := exec.Command("git", "init")
	color.Grayln(cmd.String())
	err = cmd.Run()
	if err != nil {
		color.Errorln("git init fail")
		return
	}
	return
}

func (p *TplBase) readBranchOrTagName() (name string, err error) {
	validate := func(input string) error {
		/* if len(strings.TrimSpace(input)) < 1 {
			return errors.New("the BranchName or TagName must have more than 1 characters")
		} */
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "分支 或 Tag",
		Validate: validate,
		Default:  "",
	}

	name, err = prompt.Run()

	if err != nil {
		return
	}

	name = strings.TrimSpace(name)

	return
}
