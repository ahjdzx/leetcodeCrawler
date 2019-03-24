package main

import (
	"flag"
	"fmt"
)

func main() {
	// 按需爬取题目，保存到本地文件夹中：readme.md(题目描述)，solution函数（main包）。
	// 抓取并构建卡片中的所有题目
	// 所需要的信息：
	// 简要：id,title,type,question{questionId,title,titleSlug}
	// 详细：questionId,questionTitle,categoryTitle,codeDefinition,content,translatedContent
	id := flag.String("id", "", "the ID of an problem item")
	isAll := flag.Bool("all", false, "build all problems")
	flag.Parse()
	if *isAll {
		buildAllProblems()
	} else {
		if len(*id) == 0 {
			fmt.Println("Please input the Id of problem item.")
			return
		} else {
			buildProblem(*id,"./")
		}
	}
}
