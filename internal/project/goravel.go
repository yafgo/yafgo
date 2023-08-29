package project

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gookit/color"
	"github.com/yafgo/yafgo/internal/pkg/file"
)

type TplGoravel struct {
	repo string
}

const (
	GORAVEL_REPO_GITHUB = "https://github.com/goravel/goravel.git"
)

func (rp *TplGoravel) Name() string {
	return "Goravel"
}

// MakeProject
//
//	name: like "github.com/you/project" or "my_project"
func (rp *TplGoravel) MakeProject(name string) (err error) {

	// set repo
	rp.repo = GORAVEL_REPO_GITHUB

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

	// create .env
	{
		cmd := exec.Command("cp", ".env.example", ".env")
		color.Grayln(cmd.String())
		err = cmd.Run()
		if err != nil {
			color.Errorln("create .env fail")
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

func (rp *TplGoravel) renameModule(moduleName, dir string) (err error) {

	err = file.WalkFiles(dir, func(elem string) error {
		// go.mod
		if path.Base(elem) == "go.mod" {
			err := file.ReplaceString(elem, "module goravel", "module "+moduleName)
			return err
		}
		// *.go
		if strings.HasSuffix(elem, ".go") {
			err := file.ReplaceString(elem, `"goravel/`, `"`+moduleName+`/`)
			return err
		}

		return nil
	})

	return
}
