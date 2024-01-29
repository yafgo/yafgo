package project

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gookit/color"
	"github.com/yafgo/yafgo/internal/pkg/file"
)

func NewTplYafgoLayout() *TplYafgoLayout {
	tpl := &TplYafgoLayout{
		TplBase: TplBase{
			REPO_GITHUB: "https://github.com/yafgo/yafgo-layout.git",
			REPO_GITEE:  "https://gitee.com/yafgo/yafgo-layout.git",
		},
	}
	return tpl
}

type TplYafgoLayout struct {
	TplBase
}

func (rp *TplYafgoLayout) Name() string {
	return "YafgoLayout"
}

// MakeProject
//
//	name: like "github.com/you/project" or "my_project"
func (rp *TplYafgoLayout) MakeProject(name string) (err error) {

	// select repo
	err = rp.selectRepo()
	if err != nil {
		return
	}

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
	_ = os.RemoveAll(path.Join(rp.projectDir, "LICENSE"))

	// chdir
	if _err := os.Chdir(rp.projectDir); _err != nil {
		err = _err
		return
	}

	// replace moduleName
	if err = rp.renameModule(rp.projectDir); err != nil {
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

	return
}

func (rp *TplYafgoLayout) renameModule(dir string) (err error) {

	err = file.WalkFiles(dir, func(elem string) error {
		if strings.HasSuffix(elem, ".go") || strings.HasSuffix(elem, ".gotpl") || path.Base(elem) == "go.mod" {
			// *.go, *.gotpl, go.mod
			err := file.ReplaceString(elem, "yafgo/yafgo-layout", rp.moduleName)
			return err
		}

		return nil
	})

	return
}
