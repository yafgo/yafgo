package project

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/yafgo/yafgo/internal/pkg/file"
)

func NewTplYafgoLayoutWeb() *TplYafgoLayoutWeb {
	tpl := &TplYafgoLayoutWeb{
		TplBase: TplBase{
			REPO_GITHUB: "https://github.com/yafgo/yafgo-layout-web.git",
			REPO_GITEE:  "https://gitee.com/yafgo/yafgo-layout-web.git",
		},
	}
	return tpl
}

type TplYafgoLayoutWeb struct {
	TplBase
}

func (rp *TplYafgoLayoutWeb) Name() string {
	return "YafgoLayoutWeb"
}

// MakeProject
//
//	name: like "github.com/you/project" or "my_project"
func (rp *TplYafgoLayoutWeb) MakeProject(name string) (err error) {

	// select repo
	if err = rp.selectRepo(); err != nil {
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
	_ = os.RemoveAll(filepath.Join(rp.projectDir, ".git"))
	_ = os.RemoveAll(filepath.Join(rp.projectDir, "LICENSE"))

	// chdir
	if err = os.Chdir(rp.projectDir); err != nil {
		return
	}
	// git init
	if err = rp.runGitInit(); err != nil {
		return
	}

	// go项目代码额外处理
	{
		// chdir
		dirBackend := filepath.Join(rp.projectDir, "server")
		if err = os.Chdir(dirBackend); err != nil {
			return
		}

		// replace moduleName
		if err = rp.renameModule(dirBackend); err != nil {
			return
		}

		// gofmt
		if err = rp.runGoFmt(); err != nil {
			return
		}
	}

	return
}

func (rp *TplYafgoLayoutWeb) renameModule(dir string) (err error) {

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
