package project

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gookit/color"
	"github.com/yafgo/yafgo/internal/pkg/file"
)

func NewTplGoravel() *TplGoravel {
	tpl := &TplGoravel{
		TplBase: TplBase{
			REPO_GITHUB: "https://github.com/goravel/goravel.git",
		},
	}
	return tpl
}

type TplGoravel struct {
	TplBase
}

func (rp *TplGoravel) Name() string {
	return "Goravel"
}

// MakeProject
//
//	name: like "github.com/you/project" or "my_project"
func (rp *TplGoravel) MakeProject(name string) (err error) {

	// set repo
	rp.repo = rp.REPO_GITHUB

	defer func() {
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(color.Success.Sprint("创建成功: "), rp.dirName)
	}()

	// 初始化项目信息
	if err = rp.initProjectDir(name); err != nil {
		return
	}

	// clone project
	if err = rp.cloneProject(); err != nil {
		return
	}

	// delete .git
	_ = os.RemoveAll(path.Join(rp.projectDir, ".git"))

	// replace moduleName
	rp.renameModule(rp.projectDir)

	// chdir
	if _err := os.Chdir(rp.projectDir); _err != nil {
		err = _err
		return
	}

	// gofmt
	if err = rp.runGoFmt(); err != nil {
		return
	}

	// git init
	if err = rp.runGitInit(); err != nil {
		return
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

	return
}

func (rp *TplGoravel) renameModule(dir string) (err error) {

	err = file.WalkFiles(dir, func(elem string) error {
		// go.mod
		if path.Base(elem) == "go.mod" {
			err := file.ReplaceString(elem, "module goravel", "module "+rp.moduleName)
			return err
		}
		// *.go
		if strings.HasSuffix(elem, ".go") {
			err := file.ReplaceString(elem, `"goravel/`, `"`+rp.moduleName+`/`)
			return err
		}

		return nil
	})

	return
}
