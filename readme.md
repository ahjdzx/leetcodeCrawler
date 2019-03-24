## Leetcode Crawler

用 Golang 编写的 Leetcode 题目爬取工具，可以爬取题目并按文件夹保存，方便整理自己的解题记录。

## 功能介绍

目前实现了爬取 leetcode-cn 上 golang 题目的支持，会生成 `solution.go` 模板文件和包含中文题目描述的 `readme.md`。

做这个工具是出于自用的目的，如果你有更高级需求，可以在 issue 中告诉我。

## 使用指南

### 1. 下载

如果本地已经安装 golang 环境，可以通过以下命令安装
```bash
go get github.com/Deardrops/leetcodeCrawler
go install github.com/Deardrops/leetcodeCrawler
```
你也可以从 [release 页面](https://github.com/Deardrops/leetcodeCrawler/releases) 下载对应平台的二进制文件，并手动添加到环境变量 `PATH` 中。

### 2. 运行

程序通过参数 `id` 指定题号，注意这个题号是leetcode-cn上**探索**中的题号，链接如下，url最后的 `21` 为题号。
```
https://leetcode-cn.com/explore/featured/card/top-interview-questions-easy/1/array/21/
```
调用示例：
```bash
leetcodeCrawler --id=21
```
会在当前路径下以题目名称创建新文件夹，并在新文件夹中创建题目模版和题目描述文件。

-----

也可以一键生成所有卡片中的题目，通过参数 `all` 指定。

调用示例：
```bash
leetcodeCrawler --all
```
会爬取初级算法、中级算法、高级算法三个卡片中的所有题目，并分门别类地保存到相关文件夹中。

> 如果对你有帮助的话，不妨点个 star 吧~
