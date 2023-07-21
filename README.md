# Toy - Golang Toy Cli

**Cli 工具, 用于一键创建基于内置 `go 项目模板` 的新项目**

支持创建的项目模板:

- [go-toy-layout](https://github.com/yafgo/toy-layout.git)
- [goravel](https://github.com/goravel/goravel.git)

## 环境要求:

- `git`
- `golang` _(>= 1.18)_

使用:

```shell
# 安装
go install github.com/yafgo/toy@latest

# 创建新项目
toy

# 示例
✔ Project Name: my_project
Use the arrow keys to navigate: ↓ ↑ → ←
Select Template?
  🌶 [Toy]      (Toy Layout)
     [Yafgo]    (Yafgo Layout)
     [Goravel]  (Goravel)
```
