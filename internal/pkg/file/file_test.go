package file

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	var err error
	cwd, _ := os.Getwd()
	file := cwd + "/yafgo/yafgo.txt"
	err = Create(file, []byte(`yafgo`))
	assert.NoError(t, err)

	assert.Equal(t, 1, GetLineNum(file))
	assert.True(t, Exists(file))
	assert.True(t, Remove(file))
	assert.True(t, Remove(cwd+"/yafgo"))
}

func TestPutContent(t *testing.T) {
	var err error
	cwd, _ := os.Getwd()
	file := cwd + "/yafgo/yafgo.txt"
	err = Create(file)
	assert.NoError(t, err)

	PutContent(file, []byte(`yafgo`))

	assert.Equal(t, 1, GetLineNum(file))
	assert.True(t, Exists(file))
	assert.True(t, Remove(file))
	assert.True(t, Remove(cwd+"/yafgo"))
}

func TestExtension(t *testing.T) {
	file := "path/to/file.go"
	extension := Extension(file)
	assert.NotEmpty(t, extension)
	assert.Equal(t, "go", extension)

	baseName := BaseName(file)
	assert.Equal(t, "file.go", baseName)

	baseNameWithoutExt := BaseNameWithoutExtentsion(file)
	assert.Equal(t, "file", baseNameWithoutExt)
}

func TestReplaceString(t *testing.T) {
	_file := "./demo/demo.go"
	Create(_file, []byte(`package main

import (
    "github.com/yafgo/yafgo/app"
)

const a = "hello"
`))
	ReplaceString(_file, "github.com/yafgo/yafgo", "my_project")
}

func TestWalkFiles(t *testing.T) {
	err := WalkFiles("./..", func(elem string) error {
		fmt.Println(elem)
		if path.Base(elem) == ".DS_Store" {
			return errors.New("碰到了 .DS_Store 文件")
		}
		return nil
	})
	fmt.Println(err)
}
