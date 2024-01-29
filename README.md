# Yafgo - Golang Yafgo Cli

**Cli 工具, 用于一键创建基于内置 `go 项目模板` 的新项目**

支持创建的项目模板:

- [yafgo-layout](https://github.com/yafgo/yafgo-layout.git)
- [yafgo-layout-web](https://github.com/yafgo/yafgo-layout-web.git)
- [goravel](https://github.com/goravel/goravel.git)

## 环境要求

- `git`
- `golang` _(>= 1.18)_

使用:

```shell
# 安装
go install github.com/yafgo/yafgo@latest

# 创建新项目
yafgo

# 示例
✔ Project Name: my_project
Use the arrow keys to navigate: ↓ ↑ → ←
Select Template?
  🌶 [Yafgo]    (Yafgo 后端项目模板)
     [YafgoWeb] (Yafgo 前后端项目模板)
     [Goravel]  (Goravel)
```
